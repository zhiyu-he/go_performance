package simple_impl

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestPointer(t *testing.T) {
	var a int64 = 100
	ptr := unsafe.Pointer(&a)
	fmt.Printf("%v\n", ptr)
	fmt.Printf("%v\n", &a)
}

func TestOffset(t *testing.T) {
	var m struct {
		a int
		b int
		c map[int64]int64
	}
	objPtr := &m
	s := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(objPtr)) + unsafe.Offsetof(m.b)))
	*s = 100
	fmt.Printf("%v\n", *objPtr)

}

type m struct {
	a int
	b int
	c int
}

func TestMalloc(t *testing.T) {
	var M m
	size := int(unsafe.Sizeof(M))
	fmt.Printf("sizeof(m): %d\n", size)
	for i := 0; i < 1024*24; i += 24 {
		ptr := malloc(size)
		mptr := (*m)(ptr)
		mptr.a = 100
		mptr.b = 200
		mptr.c = 200
		fmt.Printf("%v\n", *mptr)
		free(ptr, size)
	}
	ptr := malloc(24)
	fmt.Printf("%v\n", ptr)
}
