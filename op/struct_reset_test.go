package op

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"unsafe"
)

type SliceHeader struct {
	addr uintptr
	len  int
	cap  int
}

type AA struct {
	Id        int64
	name      string
	locations []string
	version   *int64
	mm        *MM
}

type MM struct {
	m1 int64
	m2 int64
}

func int64Ptr(v int64) *int64 {
	return &v
}

func TestStructAddr(t *testing.T) {
	aa := &AA{}
	ptr := unsafe.Pointer(aa)
	fmt.Printf("ptr: %v\n", ptr)
}

func TestEqual(t *testing.T) {
	a1 := &AA{
		Id:        100,
		name:      "aa",
		locations: []string{"a", "b", "c"},
		version:   int64Ptr(1),
		mm:        &MM{m1: 1, m2: 2},
	}
	a2 := &AA{
		Id:        100,
		name:      "aa",
		locations: []string{"a", "b", "c"},
		version:   int64Ptr(1),
		mm:        &MM{m1: 1, m2: 2},
	}
	a3 := &AA{
		Id:        100,
		name:      "aa",
		locations: []string{"a", "b", "c"},
		version:   int64Ptr(1),
		mm:        &MM{m1: 1, m2: 2},
	}

	empty := AA{}
	emptySize := int(unsafe.Sizeof(empty))

	Reset1(a1)
	Reset2(a2)
	Reset3(unsafe.Pointer(a3), unsafe.Pointer(&empty), emptySize)

	base := &AA{}
	assert.EqualValues(t, base, a1)
	assert.EqualValues(t, base, a2)
	assert.EqualValues(t, base, a3)

}

func BenchmarkReset1(b *testing.B) {
	a := &AA{
		Id:        100,
		name:      "aa",
		locations: []string{"a", "b", "c"},
		version:   int64Ptr(1),
		mm:        &MM{m1: 1, m2: 2},
	}
	for i := 0; i < b.N; i++ {
		Reset1(a)
	}
}

func BenchmarkReset2(b *testing.B) {
	a := &AA{
		Id:        100,
		name:      "aa",
		locations: []string{"a", "b", "c"},
		version:   int64Ptr(1),
		mm:        &MM{m1: 1, m2: 2},
	}
	for i := 0; i < b.N; i++ {
		Reset2(a)
	}
}

func BenchmarkReset3(b *testing.B) {
	a := &AA{
		Id:        100,
		name:      "aa",
		locations: []string{"a", "b", "c"},
		version:   int64Ptr(1),
		mm:        &MM{m1: 1, m2: 2},
	}
	empty := AA{}
	emptySize := int(unsafe.Sizeof(empty))

	for i := 0; i < b.N; i++ {
		Reset3(unsafe.Pointer(a), unsafe.Pointer(&empty), emptySize)
	}
}

func Reset1(v *AA) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func Reset2(v *AA) {
	v.Id = 0
	v.name = ""
	v.locations = nil
	v.version = nil
	v.mm = nil
}

func Reset3(dst unsafe.Pointer, src unsafe.Pointer, objSize int) {
	size := objSize >> 3
	a := SliceHeader{
		addr: uintptr(dst),
		len:  size,
		cap:  size,
	}
	b := SliceHeader{
		addr: uintptr(src),
		len:  size,
		cap:  size,
	}
	utb := *(*[]int64)(unsafe.Pointer(&a))
	utb1 := *(*[]int64)(unsafe.Pointer(&b))
	copy(utb, utb1)
}
