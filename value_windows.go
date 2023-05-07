package easycgo

import "C"

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

type Value struct {
	ptr uintptr
}

func (v *Value) ToBool() bool {
	return C2Go.Bool(v.ptr)
}

func (v *Value) ToString() string {
	return C2Go.String4CharPtr(v.ptr)
}

func wcharPtrToString(p *C.wchar_t) string {
	return windows.UTF16PtrToString((*uint16)(p))
}

func string4WCharPtr(p uintptr) string {
	return wcharPtrToString((*C.wchar_t)(unsafe.Pointer(p)))
}

func (v *Value) ToString4w() string {
	return string4WCharPtr(v.ptr)
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

func (v *Value) ToFloat32() float32 {
	return C2Go.Float32(v.ptr)
}

func (v *Value) ToFloat32Slice(len int) []float32 {
	return C2Go.Float32Slice(v.ptr, len)
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

func (v *Value) IsNil() bool {
	return v.ptr == 0
}

func wcharPtrFromString(s string) (*C.wchar_t, error) {
	p, err := windows.UTF16PtrFromString(s)
	return (*C.wchar_t)(p), err
}

func wcharUint16FromString(s string) (*uint16, error) {
	p, err := windows.UTF16PtrFromString(s)
	return p, err
}

func WChars(d string) uintptr {
	sp, _ := wcharPtrFromString(d)
	return uintptr(unsafe.Pointer(&sp))
}

func WChars2Uint16(d string) *uint16 {
	sp, _ := wcharUint16FromString(d)
	return sp
}
