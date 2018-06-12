package simple_impl

import "testing"

type M struct {
	Q1 int64
	Q2 int64
	Q3 int64
}

type A struct {
	Id       int64
	Name     string
	Ptr      *M
	NoUsePtr *M
	Sli1     []int64
	Sli2     []*M
	Arr3     [3]int
}

type A2 struct {
	M1 map[int]int
	M2 map[string]*M
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

	a2 := &A2{}
	MeasureObject(a2)

	a3 := &A2{M1: nil, M2: map[string]*M{"t1": m, "t2": m}}
	MeasureObject(a3)
}
