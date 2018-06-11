package simple_impl

import "testing"

type M struct {
	Q1 int64
	Q2 int64
	Q3 int64
}

type A struct {
	Id   int64
	Name string
	Ptr  *M
	Sli1 []int64
	Sli2 []*M
	Arr3 [3]int
}

func TestMeasureObject(t *testing.T) {
	m := &M{
		Q1: 101,
		Q2: 102,
		Q3: 103,
	}
	sli1 := []int64{1, 2, 3}
	sli2 := make([]*M, 10)
	a := &A{Id: 100, Name: "toutiao", Ptr: m, Sli1: sli1, Sli2: sli2}
	MeasureObject(a)
}
