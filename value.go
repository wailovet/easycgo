package easycgo

import "os"

type ValueInf interface {
	SetPtr(ptr uintptr)
	ToBool() bool
	ToString() string
	ToString4w() string
	ToStringSlice(len int) []string
	ToFloat64() float64
	ToFloat64Slice(len int) []float64
	ToFloat32() float32
	ToFloat32Slice(len int) []float32
	ToInt64() int64
	ToInt64Slice(len int) []int64
	ToInt() int
	ToIntSlice(len int) []int
	ToUint() uint
	ToUintSlice(len int) []uint
	ToUint64() uint64
	ToUint64Slice(len int) []uint64
	ToUint32() uint32
	ToUint32Slice(len int) []uint32
	ToUint8() uint8
	ToUint8Slice(len int) []uint8
	ToBytes(len int) []byte
	ToBytes4Char() []byte
	Value() interface{}
	IsNil() bool
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
