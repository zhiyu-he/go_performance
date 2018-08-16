package simple_impl

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

var objPool *ObjectPool
var objChan chan map[int64]int64

func init() {
	f := func() interface{} {
		return make(map[int64]int64, 24)
	}
	objPool = NewObjectPool("map", 1000000, f)

	objChan = make(chan map[int64]int64, 1000000)

	for i := 0; i < 1000000; i++ {
		objChan <- make(map[int64]int64, 24)
	}
}

func GetOBJ() {
	for i := 0; i < 200000; i++ {
		b := objPool.Borrow()
		objPool.Return(b)
	}
}

func GetOBJ2() {
	for i := 0; i < 200000; i++ {
		_ = make(map[int64]int64, 24)
	}
}

func GetOBJ3() {
	for i := 0; i < 200000; i++ {
		b := <-objChan
		objChan <- b
	}
}

func BenchmarkGetOBJ(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetOBJ()
	}
}

func BenchmarkGetOBJ2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetOBJ2()
	}
}

func BenchmarkGetOBJ3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetOBJ3()
	}
}

func TestUnsortedInsert(t *testing.T) {
	c := 100
	pq := newMinHeap(c)
	ints := make([]int, 0, c)
	for i := 0; i < c; i++ {
		v := rand.Int()
		ints = append(ints, v)
		heap.Push(&pq, &Dirty{item: nil, ts: int64(v)})
	}
	assert.Equal(t, pq.Len(), c)
	assert.Equal(t, cap(pq), c)
	sort.Sort(sort.IntSlice(ints))
	for i := 0; i < c; i++ {
		item, _ := pq.PeekAndShift(int64(ints[len(ints)-1]))
		assert.Equal(t, item.ts, int64(ints[i]))
	}
}
