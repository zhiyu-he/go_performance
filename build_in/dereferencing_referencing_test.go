package build_in

import "testing"

type O struct {
	Num int32
	NumPtr *int32
}

var Obj *O

func init() {
	var num int32 = 100

	Obj = &O{
		Num: 100,
		NumPtr: &num,
	}
}


func ReadByReferencing() int32 {
	return Obj.Num
}

func ReadByDereferencing() int32 {
	return *Obj.NumPtr
}


func BenchmarkReadByReferencing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadByReferencing()
	}
}

func BenchmarkReadByDereferencing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadByDereferencing()
	}
}
