package easycgo

import "syscall"

type EasyCgo struct {
	dll *syscall.DLL
}
type EasyCgoProc struct {
	proc *syscall.Proc
}

func MustLoad(filename string) *EasyCgo {
	return &EasyCgo{
		dll: syscall.MustLoadDLL(filename),
	}
}

func Load(filename string) (*EasyCgo, error) {
	dll, err := syscall.LoadDLL(filename)
	if err != nil {
		return nil, err
	}
	return &EasyCgo{
		dll: dll,
	}, nil
}

func (ec *EasyCgo) Find(name string) (*EasyCgoProc, error) {
	proc, err := ec.dll.FindProc(name)
	return &EasyCgoProc{
		proc: proc,
	}, err
}

func (ec *EasyCgo) MustFind(name string) *EasyCgoProc {
	return &EasyCgoProc{
		proc: ec.dll.MustFindProc(name),
	}
}

func (ecp *EasyCgoProc) Call(args ...interface{}) ValueInf {

	uiargs := []uintptr{}
	for i := range args {
		uiargs = append(uiargs, Go2C.Auto(args[i]))
	}
	r1, _, _ := ecp.proc.Call(uiargs...)
	return &Value{
		ptr: r1,
	}
}
