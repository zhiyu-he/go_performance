package simple_impl

import (
	"testing"
	//"github.com/stretchr/testify/assert"
	"unsafe"
)

const (
	pool_size = 600
	slice_cap = 8192
)
var pool = NewSimpleSlicePool(pool_size, slice_cap)

/*
func TestSimpleSlicePool_GetSliceObject(t *testing.T) {
	var lastOne *[]interface{}
	var lastIdx int
	for i := 0; i < pool_size-1; i++ {
		lastOne, lastIdx = pool.GetSliceObject()
		assert.NotNil(t, lastOne)
		assert.Equal(t, i, lastIdx)
	}
	// 执行第600次, 因为没有put操作, 因此会返回(nil, -1)
	obj, idx := pool.GetSliceObject()
	assert.Nil(t, obj)
	assert.Equal(t, -1, idx)

	// return一个
	m := (*[]*int)(unsafe.Pointer(lastOne))
	one := 1
	two := 2
	*m = append(*m, &one)
	*m = append(*m, &two)
	pool.PutSliceObject(lastOne, lastIdx)

	// 再拿一个出来
	obj, idx = pool.GetSliceObject()
	assert.NotNil(t, obj)
	assert.Equal(t, 599, idx)


	obj, idx = pool.GetSliceObject()
	assert.Nil(t, obj)
	assert.Equal(t, -1, idx)
}
*/

type TestObj struct {
	Index int
}

func doWithSlicePool() {
	obj, idx := pool.GetSliceObject()

	m := (*[]*TestObj)(unsafe.Pointer(obj))

	for i := 0;  i < 8192; i++ {
		*m = append(*m, &TestObj{Index: i})
	}

	pool.PutSliceObject(obj, idx)
}

func doWithSliceMake() {
	m := make([]*TestObj, 0, 8192)
	for i := 0;  i < 8192; i++ {
		m = append(m, &TestObj{Index: i})
	}
}




func BenchmarkDoWithSlicePool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doWithSlicePool()
	}
}

func BenchmarkDoWithSliceMake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doWithSliceMake()
	}
}
