package easycgo

import "C"

import (
	"fmt"
	"unsafe"
)

var Go2C Go2CBase

type Go2CBase struct{}

func (g2c *Go2CBase) AnyPtr(d interface{}) uintptr {
	return uintptr(unsafe.Pointer(&d))
}

func (g2c *Go2CBase) Uintptr(d interface{}) uintptr {
	return d.(uintptr)
}

func (g2c *Go2CBase) IntSlice(d []int) uintptr {
	if d == nil {
		return 0
	}
	ret := unsafe.Pointer(&d[0])
	return uintptr(unsafe.Pointer(ret))
}

func (g2c *Go2CBase) Int8Slice(d []int8) uintptr {
	if d == nil {
		return 0
	}
	ret := unsafe.Pointer(&d[0])
	return uintptr(unsafe.Pointer(ret))
}

func (g2c *Go2CBase) Int16Slice(d []int16) uintptr {
	if d == nil {
		return 0
	}
	ret := unsafe.Pointer(&d[0])
	return uintptr(unsafe.Pointer(ret))
}

func (g2c *Go2CBase) Int32Slice(d []int32) uintptr {
	if d == nil {
		return 0
	}
	ret := unsafe.Pointer(&d[0])
	return uintptr(unsafe.Pointer(ret))
}

func (g2c *Go2CBase) Int64Slice(d []int64) uintptr {
	if d == nil {
		return 0
	}
	ret := unsafe.Pointer(&d[0])
	return uintptr(unsafe.Pointer(ret))
}

func (g2c *Go2CBase) UintSlice(d []uint) uintptr {
	if d == nil {
		return 0
	}
	ret := unsafe.Pointer(&d[0])
	return uintptr(unsafe.Pointer(ret))
}

func (g2c *Go2CBase) Uint8Slice(d []uint8) uintptr {
	if d == nil {
		return 0
	}
	ret := unsafe.Pointer(&d[0])
	return uintptr(unsafe.Pointer(ret))
}

func (g2c *Go2CBase) Uint16Slice(d []uint16) uintptr {
	if d == nil {
		return 0
	}
	ret := unsafe.Pointer(&d[0])
	return uintptr(unsafe.Pointer(ret))
}

func (g2c *Go2CBase) Uint32Slice(d []uint32) uintptr {
	if d == nil {
		return 0
	}
	ret := unsafe.Pointer(&d[0])
	return uintptr(unsafe.Pointer(ret))
}

func (g2c *Go2CBase) Uint64Slice(d []uint64) uintptr {
	if d == nil {
		return 0
	}
	ret := unsafe.Pointer(&d[0])
	return uintptr(unsafe.Pointer(ret))
}

func (g2c *Go2CBase) UintPtrSlice(d []uintptr) uintptr {
	if d == nil {
		return 0
	}
	ret := unsafe.Pointer(&d[0])
	return uintptr(unsafe.Pointer(ret))
}

func (g2c *Go2CBase) Float64Slice(d []float64) uintptr {
	if d == nil {
		return 0
	}
	ret := unsafe.Pointer(&d[0])
	return uintptr(ret)
}

func (g2c *Go2CBase) Float32Slice(d []float32) uintptr {
	if d == nil {
		return 0
	}
	ret := unsafe.Pointer(&d[0])
	return uintptr(ret)
}

func (g2c *Go2CBase) Int(d int) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Int8(d int8) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Int16(d int16) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Int32(d int32) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Int64(d int64) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Float32(d float32) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Float64(d float64) uintptr {
	// return uintptr(math.Float64bits(d)) //?
	return uintptr(d)
}

func (g2c *Go2CBase) Byte(d byte) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Uint8(d uint8) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Uint16(d uint16) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Uint(d uint) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Uint32(d uint32) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Uint64(d uint64) uintptr {
	return uintptr(d)
}

func (g2c *Go2CBase) Bool(d bool) uintptr {
	if d {
		return 1
	} else {
		return 0
	}
}

func (g2c *Go2CBase) Chars(d string) uintptr {
	b := []byte(d)
	var newByte []byte
	for i := range b {
		newByte = append(newByte, b[i])
	}
	newByte = append(newByte, 0)
	// log.Println("Chars", b)
	return uintptr(unsafe.Pointer(&newByte[0]))
}

func (g2c *Go2CBase) Auto(vt interface{}) (p uintptr) {
	switch v := vt.(type) {
	case string:
		p = Go2C.Chars(v)
	case []byte:
		p = Go2C.Chars(string(v))
	case bool:
		p = Go2C.Bool(v)
	case int:
		p = Go2C.Int(v)
	case int8:
		p = Go2C.Int8(v)
	case int16:
		p = Go2C.Int16(v)
	case int32:
		p = Go2C.Int32(v)
	case int64:
		p = Go2C.Int64(v)
	case uint:
		p = Go2C.Uint(v)
	case uint8:
		p = Go2C.Uint8(v)
	case uint16:
		p = Go2C.Uint16(v)
	case uint32:
		p = Go2C.Uint32(v)
	case uint64:
		p = Go2C.Uint64(v)
	case float64:
		p = Go2C.Float64(v)
	case float32:
		p = Go2C.Float32(v)
	case []float32:
		p = Go2C.Float32Slice(v)
	case []float64:
		p = Go2C.Float64Slice(v)
	case []uint:
		p = Go2C.UintSlice(v)
	case []uint16:
		p = Go2C.Uint16Slice(v)
	case []uint32:
		p = Go2C.Uint32Slice(v)
	case []uint64:
		p = Go2C.Uint64Slice(v)
	case []int:
		p = Go2C.IntSlice(v)
	case []int16:
		p = Go2C.Int16Slice(v)
	case []int32:
		p = Go2C.Int32Slice(v)
	case []int64:
		p = Go2C.Int64Slice(v)
	case []uintptr:
		p = Go2C.UintPtrSlice(v)
	case uintptr:
		p = Go2C.Uintptr(v)
	case nil:
		p = uintptr(unsafe.Pointer(nil))
	case ValueInf:
		p = v.Value().(uintptr)
	default:
		panic(fmt.Sprintln("type error:", v))
	}
	return
}
