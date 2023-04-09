package easycgo

import (
	"bufio"
	"debug/elf"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func CheckErrorWithLDD(pathname string) {
	if runtime.GOOS != "linux" {
		return
	}
	data, err := LDD(pathname)
	if err != nil {
		panic(err.Error())
	}
	maxLen := 0
	emptySO := []string{}
	for i := range data {
		if len([]rune(i)) > maxLen {
			maxLen = len([]rune(i))
		}
		if data[i] == "" {
			emptySO = append(emptySO, i)
		}
	}

	if len(emptySO) > 0 {
		fmt.Println("缺少动态库: ", strings.Join(emptySO, " , "))
		fmt.Println("---------------")

		for i := range data {
			fmt.Printf(fmt.Sprint("%-", maxLen, "s => %s \n"), i, data[i])
		}

		dirstr, _ := filepath.Abs(filepath.Dir(pathname))
		fmt.Println(`尝试 export LD_LIBRARY_PATH="$LD_LIBRARY_PATH;` + dirstr + `"`)
		os.Exit(1)
	}
}

func LDD(pathname string) (map[string]string, error) {
	if runtime.GOOS != "linux" {
		return nil, errors.New("仅支持")
	}
	m := &lddT{
		depsStr: map[string]string{},
	}

	m.deps = make(map[string]depsInfo)
	m.deflib = []string{"/lib/", "/usr/lib/", "/lib64", "/usr/lib64"}
	m.envlib = os.Getenv("LD_LIBRARY_PATH")
	m.conflib = m.readLdSoConf("/etc/ld.so.conf", m.conflib)

	pathname = m.realPath(m.findLib(pathname, nil))
	// log.Println("pathname:", pathname)
	f, err := elf.Open(pathname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	m.deps_root = new(depsNode)
	m.deps_root.name = path.Base(pathname)

	m.deps_list = append(m.deps_list, m.deps_root)
	for len(m.deps_list) > 0 {
		// pop first element
		dep := m.deps_list[0]
		m.deps_list = m.deps_list[1:]

		m.processDep(pathname, dep)
	}
	m.depTree(m.deps_root, f)
	return m.depsStr, nil
}

func (m *lddT) depTree(n *depsNode, f *elf.File) {
	m.depsStr[n.name] = m.deps[n.name].path

	for _, v := range n.child {
		m.depTree(v, f)
	}
}

func (m *lddT) printDepTree(n *depsNode, f *elf.File) {
	for i := 0; i < n.depth; i++ {
		fmt.Printf("   ")
	}

	fmt.Printf("%s  => %s\n", n.name, m.deps[n.name].path)

	for _, v := range n.child {
		m.printDepTree(v, f)
	}
}

func (m *lddT) readLdSoConf(name string, libpath []string) []string {
	// log.Println("readLdSoConf:", name)
	f, err := os.Open(name)
	if err != nil {
		return libpath
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		t := s.Text()

		if len(strings.TrimSpace(t)) == 0 {
			continue
		}
		if strings.HasPrefix(t, "#") {
			continue
		}

		if strings.HasPrefix(t, "include") {
			libs, err := filepath.Glob(t[8:])
			if err != nil {
				continue
			}
			for _, l := range libs {
				libpath = m.readLdSoConf(l, libpath)
			}
		} else {
			libpath = append(libpath, t)
		}
	}
	return libpath
}

type depsInfo struct {
	path   string
	mach   elf.Machine
	bits   elf.Class
	endian binary.ByteOrder
	kind   elf.Type
	abi    elf.OSABI
	ver    uint8

	libs []string
	isym []elf.ImportedSymbol
	dsym []elf.Symbol
	syms []elf.Symbol
	prog []*elf.Prog
	sect []*elf.Section
	dyns []dynInfo
}

type depsNode struct {
	name   string
	parent *depsNode
	child  []*depsNode
	depth  int
}
type dynInfo struct {
	tag elf.DynTag
	val interface{}
}

type lddT struct {
	deps      map[string]depsInfo
	deps_list []*depsNode
	deps_root *depsNode
	deflib    []string
	envlib    string
	conflib   []string
	depsStr   map[string]string
}

// search shared libraries as described in `man ld.so(8)`
func (m *lddT) findLib(name string, parent *depsNode) string {
	if strings.Contains(name, "/") {
		return name
	}

	// check DT_RPATH attribute
	if parent != nil {
		info := m.deps[parent.name]
		for _, dyn := range info.dyns {
			if dyn.tag != elf.DT_RPATH {
				continue
			}

			fullpath := path.Join(dyn.val.(string), name)
			// log.Println("fullpath 1:", fullpath)
			if _, err := os.Stat(fullpath); err == nil {
				return fullpath
			}
		}
	}

	m.envlib = strings.ReplaceAll(m.envlib, ":", ";")
	// check LD_LIBRARY_PATH environ
	for _, libpath := range strings.Split(m.envlib, ";") {
		fullpath := path.Join(libpath, name)
		// log.Println("fullpath 2:", fullpath)
		if _, err := os.Stat(fullpath); err == nil {
			return fullpath
		}
	}

	// check DT_RUNPATH attribute
	if parent != nil {
		info := m.deps[parent.name]
		for _, dyn := range info.dyns {
			if dyn.tag != elf.DT_RUNPATH {
				continue
			}

			fullpath := path.Join(dyn.val.(string), name)
			// log.Println("fullpath 3:", fullpath)
			if _, err := os.Stat(fullpath); err == nil {
				return fullpath
			}
		}
	}

	// check libraries in /etc/ld.so.conf
	for _, libpath := range m.conflib {
		fullpath := path.Join(libpath, name)
		// log.Println("fullpath 4:", fullpath)
		if _, err := os.Stat(fullpath); err == nil {
			return fullpath
		}
	}

	// check default library directories
	for _, libpath := range m.deflib {
		fullpath := path.Join(libpath, name)
		// log.Println("fullpath 5:", fullpath)
		if _, err := os.Stat(fullpath); err == nil {
			return fullpath
		}
	}
	return ""
}

func (m *lddT) realPath(pathname string) string {
	if pathname == "" {
		return ""
	}

	relpath, _ := filepath.EvalSymlinks(pathname)
	abspath, _ := filepath.Abs(relpath)

	return abspath
}
func (m *lddT) processDep(pathname string, dep *depsNode) error {
	// skip duplicate libraries
	if _, ok := m.deps[dep.name]; ok {
		return nil
	}

	info := depsInfo{path: m.realPath(m.findLib(dep.name, dep.parent))}

	if dep.parent == nil {
		info.path = m.realPath(pathname)
	}

	f, err := elf.Open(info.path)
	if err != nil {
		return fmt.Errorf("%v: %s (%s)\n", err, info.path, dep.name)

	}
	defer f.Close()

	info.mach = f.Machine
	info.bits = f.Class
	info.kind = f.Type
	info.abi = f.OSABI
	info.ver = f.ABIVersion
	info.endian = f.ByteOrder

	info.prog = f.Progs
	info.sect = f.Sections

	if f.Type != elf.ET_EXEC && f.Type != elf.ET_DYN {
		return fmt.Errorf("elftree: `%s` seems not to be a valid ELF executable\n", dep.name)
	}

	if m.readDynamic(f, &info) < 0 {
		return fmt.Errorf("elftree: `%s` seems to be statically linked\n", dep.name)
	}

	libs, err := f.ImportedLibraries()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	isym, err := f.ImportedSymbols()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	dsym, err := f.DynamicSymbols()
	if err != nil {
		return fmt.Errorf(err.Error())

	}

	syms, err := f.Symbols()
	if err == nil {
		info.syms = syms
	}

	info.libs = libs
	info.dsym = dsym
	info.isym = isym

	var L []*depsNode
	for _, soname := range libs {
		N := new(depsNode)
		N.name = soname
		N.parent = dep
		N.depth = dep.depth + 1

		L = append(L, N)
		dep.child = append(dep.child, N)
	}

	m.deps_list = append(L, m.deps_list...)
	m.deps[dep.name] = info
	return nil
}

func (m *lddT) readDynamic(f *elf.File, info *depsInfo) int {
	var i, count uint

	dyn := f.Section(".dynamic")
	if dyn == nil {
		return -1
	}

	data, err := dyn.Data()
	if err != nil {
		return -1
	}
	str := f.Section(".dynstr")
	stab, err := str.Data()
	if err != nil {
		return -1
	}

	count = uint(dyn.Size / dyn.Entsize)
	for i = 0; i < count; i++ {
		var tag, val uint64

		if f.Class == elf.ELFCLASS64 {
			tag = f.ByteOrder.Uint64(data[(i*2+0)*8 : (i*2+1)*8])
			val = f.ByteOrder.Uint64(data[(i*2+1)*8 : (i*2+2)*8])
		} else {
			tag = uint64(f.ByteOrder.Uint32(data[(i*2+0)*4 : (i*2+1)*4]))
			val = uint64(f.ByteOrder.Uint32(data[(i*2+1)*4 : (i*2+2)*4]))
		}

		dtag := elf.DynTag(tag)
		switch dtag {
		case elf.DT_NULL:
			break
		case elf.DT_NEEDED:
			fallthrough
		case elf.DT_RPATH:
			fallthrough
		case elf.DT_RUNPATH:
			fallthrough
		case elf.DT_SONAME:
			sval := m.readElfString(stab, val)
			info.dyns = append(info.dyns, dynInfo{dtag, sval})
			break
		default:
			info.dyns = append(info.dyns, dynInfo{dtag, val})
			break
		}
	}
	return 0
}

func (m *lddT) readElfString(strtab []byte, i uint64) string {
	var len uint64

	for len = 0; strtab[i+len] != '\x00'; len++ {
		continue
	}

	return string(strtab[i : i+len])
}
