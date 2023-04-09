package easycgo

import "C"

type Value struct {
	ptr uintptr
}

func (v *Value) ToBool() bool {
	return C2Go.Bool(v.ptr)
}

func (v *Value) ToString() string {
	return C2Go.String4CharPtr(v.ptr)
}

func (v *Value) ToString4w() string {
	return v.ToString()
}

func (v *Value) ToStringSlice(len int) []string {
	panic("not implemented") // TODO: Implement
}

func (v *Value) ToFloat64() float64 {
	return C2Go.Float64(v.ptr)
}

func (v *Value) ToFloat64Slice(len int) []float64 {
	return C2Go.Float64Slice(v.ptr, len)
}

func (v *Value) ToInt64() int64 {
	return C2Go.Int64(v.ptr)
}

func (v *Value) ToInt64Slice(len int) []int64 {
	return C2Go.Int64Slice(v.ptr, len)
}

func (v *Value) ToInt() int {
	return C2Go.Int(v.ptr)
}

func (v *Value) ToIntSlice(len int) []int {
	return C2Go.IntSlice(v.ptr, len)
}

func (v *Value) ToUint() uint {
	return C2Go.Uint(v.ptr)
}

func (v *Value) ToUintSlice(len int) []uint {
	return C2Go.UintSlice(v.ptr, len)
}

func (v *Value) ToUint64() uint64 {
	return C2Go.Uint64(v.ptr)
}

func (v *Value) ToUint64Slice(len int) []uint64 {
	return C2Go.Uint64Slice(v.ptr, len)
}

func (v *Value) ToUint32() uint32 {
	return C2Go.Uint32(v.ptr)
}

func (v *Value) ToUint32Slice(len int) []uint32 {
	return C2Go.Uint32Slice(v.ptr, len)
}

func (v *Value) ToUint8() uint8 {
	return C2Go.Uint8(v.ptr)
}

func (v *Value) ToUint8Slice(len int) []uint8 {
	return C2Go.Uint8Slice(v.ptr, len)
}

func (v *Value) ToBytes(len int) []byte {
	return C2Go.Bytes4Void(v.ptr, len)
}

func (v *Value) ToBytes4Char() []byte {
	return C2Go.Bytes(v.ptr)
}

func (v *Value) Value() interface{} {
	return v.ptr
}
