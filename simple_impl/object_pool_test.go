package simple_impl

import (
	"testing"
	"time"
	"fmt"
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
		objChan<-make(map[int64]int64, 24)
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
		_=make(map[int64]int64, 24)
	}
}


func GetOBJ3() {
	for i := 0; i < 200000; i++ {
		b := <-objChan
		objChan<-b
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


func TestMinHeap(t *testing.T) {
	minHeap := newMinHeap(100)

	dirty1 := &Dirty{
		ts: time.Now().UnixNano(),
		item: make(map[int64]int64),
	}
	fmt.Printf("1: %+v", *dirty1)
	minHeap.Push(dirty1)
	time.Sleep(1 * time.Second)
	dirty2 := &Dirty{
		ts: time.Now().UnixNano(),
		item: make(map[int64]int64),
	}
	fmt.Printf("2: %+v", *dirty2)
	minHeap.Push(dirty2)

	item, ts := minHeap.PeekAndShift(time.Now().UnixNano() - int64(time.Second))
	fmt.Printf("ans: %+v %d\n", *item, ts)
}
