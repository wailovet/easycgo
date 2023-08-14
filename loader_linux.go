package easycgo

/*
#include <dlfcn.h>
#include <stdlib.h>
void* Syscall(void *handle,char *symbol,void *arg1,void *arg2,void *arg3) {
	void *(*pFunc)(void *arg1,void *arg2,void *arg3);
	pFunc = (void * (*)(void *,void *,void *))dlsym(handle,symbol);
	return (*pFunc)(arg1,arg2,arg3);
}

void* Syscall6(void *handle,char *symbol,void *arg1,void *arg2,void *arg3,void *arg4,void *arg5,void *arg6) {
	void *(*pFunc)(void *,void *,void *,void *,void *,void *);
	pFunc = (void * (*)(void *,void *,void *,void *,void *,void *))dlsym(handle,symbol);
	return (*pFunc)(arg1,arg2,arg3,arg4,arg5,arg6);
}

void* Syscall12(void *handle,char *symbol,void *arg1,void *arg2,void *arg3,void *arg4,void *arg5,void *arg6,void *arg7,void *arg8,void *arg9,void *arg10,void *arg11,void *arg12) {
	void *(*pFunc)(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *);
	pFunc = (void * (*)(void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *,void *))dlsym(handle,symbol);
	return (*pFunc)(arg1,arg2,arg3,arg4,arg5,arg6,arg7,arg8,arg9,arg10,arg11,arg12);
}
*/
// #cgo LDFLAGS: -ldl
import "C"

import (
	"fmt"
	"unsafe"
)

type EasyCgo struct {
	so uintptr
}
type EasyCgoProc struct {
	ec      *EasyCgo
	funName string
}

func MustLoad(filename string) *EasyCgo {
	CheckErrorWithLDD(filename)

	so := C.dlopen(C.CString(filename), C.RTLD_NOW)

	if so == nil {
		cerr := C.dlerror()
		panic(fmt.Sprintf("dlopen error: %s", C.GoString(cerr)))
	}

	return &EasyCgo{
		so: uintptr(so),
	}
}

func Load(filename string) (*EasyCgo, error) {
	CheckErrorWithLDD(filename)

	fmt.Println(filename, " 加载中...")
	if !pathExists(filename) {
		return nil, fmt.Errorf("file not exists:", filename)
	}
	so := C.dlopen(C.CString(filename), C.RTLD_NOW)
	if so == nil {
		cerr := C.dlerror()
		return nil, fmt.Errorf("dlopen error: %s", C.GoString(cerr))
	}
	return &EasyCgo{
		so: uintptr(so),
	}, nil
}

func (ec *EasyCgo) Find(name string) (*EasyCgoProc, error) {
	return &EasyCgoProc{
		funName: name,
		ec:      ec,
	}, nil
}

func (ec *EasyCgo) MustFind(name string) *EasyCgoProc {
	return &EasyCgoProc{
		funName: name,
		ec:      ec,
	}
}

func (ec *EasyCgo) Release() {
	C.dlclose(unsafe.Pointer(ec.so))
	ec.so = 0
}

var null = unsafe.Pointer(nil)

func (ecp *EasyCgoProc) call(args ...interface{}) uintptr {
	a := []unsafe.Pointer{}
	for i := range args {
		a = append(a, unsafe.Pointer(Go2C.Auto(args[i])))
	}
	switch len(a) {
	case 0:
		return (uintptr)(C.Syscall(unsafe.Pointer(ecp.ec.so), C.CString(ecp.funName), null, null, null))
	case 1:
		return (uintptr)(C.Syscall(unsafe.Pointer(ecp.ec.so), C.CString(ecp.funName), a[0], null, null))
	case 2:
		return (uintptr)(C.Syscall(unsafe.Pointer(ecp.ec.so), C.CString(ecp.funName), a[0], a[1], null))
	case 3:
		return (uintptr)(C.Syscall(unsafe.Pointer(ecp.ec.so), C.CString(ecp.funName), a[0], a[1], a[2]))
	case 4:
		return (uintptr)(C.Syscall6(unsafe.Pointer(ecp.ec.so), C.CString(ecp.funName), a[0], a[1], a[2], a[3], null, null))
	case 5:
		return (uintptr)(C.Syscall6(unsafe.Pointer(ecp.ec.so), C.CString(ecp.funName), a[0], a[1], a[2], a[3], a[4], null))
	case 6:
		return (uintptr)(C.Syscall6(unsafe.Pointer(ecp.ec.so), C.CString(ecp.funName), a[0], a[1], a[2], a[3], a[4], a[5]))
	case 7:
		return (uintptr)(C.Syscall12(unsafe.Pointer(ecp.ec.so), C.CString(ecp.funName), a[0], a[1], a[2], a[3], a[4], a[5], a[6], null, null, null, null, null))
	case 8:
		return (uintptr)(C.Syscall12(unsafe.Pointer(ecp.ec.so), C.CString(ecp.funName), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], null, null, null, null))
	case 9:
		return (uintptr)(C.Syscall12(unsafe.Pointer(ecp.ec.so), C.CString(ecp.funName), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], null, null, null))
	case 10:
		return (uintptr)(C.Syscall12(unsafe.Pointer(ecp.ec.so), C.CString(ecp.funName), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], null, null))
	case 11:
		return (uintptr)(C.Syscall12(unsafe.Pointer(ecp.ec.so), C.CString(ecp.funName), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], null))
	case 12:
		return (uintptr)(C.Syscall12(unsafe.Pointer(ecp.ec.so), C.CString(ecp.funName), a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11]))
	default:
		panic("Call " + ecp.funName + " with too many arguments")
	}
}

func (ecp *EasyCgoProc) Call(args ...interface{}) ValueInf {
	r1 := ecp.call(args...)
	return &Value{
		ptr: r1,
	}
}
