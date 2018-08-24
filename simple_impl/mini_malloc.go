package simple_impl

import (
	"reflect"
	"unsafe"
)

/*
	Q1: uintptr和Pointer的使用场景?
	A1:
*/

var alloc []byte
var offset int

func init() {
	alloc = make([]byte, 1024*24) // 10M
}

func malloc(size int) unsafe.Pointer {
	if offset+size > cap(alloc) {
		return nil
	}
	b := alloc[offset : offset+size]
	offset += size
	header := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	return unsafe.Pointer(header.Data)
}

func free(addr unsafe.Pointer, size int) {
	offset -= size
}
