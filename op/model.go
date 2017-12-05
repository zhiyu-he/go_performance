package op

import "unsafe"

type M struct {
	num int64
}

type BidModel struct {
	field0  int64
	field1  int64
	field3  int64
	field4  int64
	field5  int64
	field6  int64
	field7  int64
	field8  int64
	field9  int64
	field10 int64
	field11 int64
	field12 int64
	field13 int64
	field14 int64
	field15 int64
	field16 int64
	field17 int64
	field18 int64
	field19 int64
	field20 int64
	field21 int64
	field22 int64
	field23 int64
	field24 int64
	field25 int64
	field26 int64
	field27 int64
	field28 int64
	field29 int64
	field30 int64
	field31 int64
	field32 int64
	m1      *M
	m2      *M
}

func (p *BidModel) GetValue(offset uintptr) int64 {
	return *(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + offset))
}

func (p *BidModel) SetValue(offset uintptr, value int64) {
	*(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + offset)) = value
}

func (p *BidModel) SetStructValue(offset uintptr, val *M) {
	*(*M)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + offset)) = *val
}

func (p *BidModel) GetStructValue(offset uintptr) *M {
	return (*M)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + offset))
}
