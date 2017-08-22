package build_in

import (
	"encoding/binary"
	"testing"
	"unsafe"
	//"fmt"
)

type OurSlice struct {
	ptr unsafe.Pointer
	len int
	cap int
}

const (
	length_int32 = 4
	length_int64 = 8
)

var local_int_length = int(unsafe.Sizeof(int(0)))

func TestLocalIntLength(t *testing.T) {
	t.Logf("LocalIntLength: %v\n", local_int_length)
}

func TestGoSliceMemLayout(t *testing.T) {
	s := make([]int64, 0, 20)
	t.Logf("slice-values: %v\n", s)
	t.Logf("slice-Addr: %v\n", unsafe.Pointer(&s)) // slice自己的地址, 这个地址是slice-struct的地址, 里面data的地址如何表示呢？

	sHeader := *(*[24]byte)(unsafe.Pointer(&s)) // slice在内存中的元素
	t.Logf("slice-memory-layout: %v\n", sHeader)
	t.Logf("slice-data-ptr: 0x%x\n", binary.LittleEndian.Uint64(sHeader[0:8]))
	t.Logf("slice-len: %v\n", binary.LittleEndian.Uint64(sHeader[8:16]))
	t.Logf("slice-cap: %v\n", int64(binary.LittleEndian.Uint64(sHeader[16:])))

	s = append(s, 110)
	s = append(s, 119)
	s = append(s, 120)
	/*
		s[0] = 110
		s[1] = 119
		s[2] = 120
	*/

	ourSlice := (*OurSlice)(unsafe.Pointer(&s))
	t.Logf("our-slice-data-ptr: %v\n", ourSlice.ptr)
	t.Logf("our-slice-len: %v\n", ourSlice.len)
	t.Logf("our-slice-cap: %v\n", ourSlice.cap)

	/*
		tmp := &OurSlice{
			ptr: ourSlice.ptr,
			len: ourSlice.len,
			cap: ourSlice.len,
		}
	*/

	t.Logf("our-slice-values: %v\n", *(*[]int64)(unsafe.Pointer(ourSlice)))
}

var commonPool = make([]interface{}, 0, 8192)

var commonPoolAddr = &commonPool

func TestSlicePool(t *testing.T) {
	commonPool := make([]interface{}, 0, 1000)
	commonPoolAddr := (*[]interface{})(unsafe.Pointer(&commonPool))

	realSlice := (*[]int64)(unsafe.Pointer(commonPoolAddr))
	t.Logf("common-pool-addr: %v\n", unsafe.Pointer(&commonPool))
	t.Logf("real-slice-addr: %v\n", unsafe.Pointer(realSlice))

	*realSlice = append(*realSlice, 1)
	*realSlice = append(*realSlice, 2)
	*realSlice = append(*realSlice, 3)

	t.Logf("%#v\n", *realSlice)
}

type SimpleObject struct {
	Index int64
}

func GetSlice() *[]interface{} {
	return commonPoolAddr
}

func ResetSlice() {
	*commonPoolAddr = (*commonPoolAddr)[:0]
}

func doWithPool() {

	//fmt.Printf("before: len: %d cap: %d\n", len(commonPool), cap(commonPool))
	//realSlice := (*[]*SimpleObject)(unsafe.Pointer(GetSlice()))
	_ = (*[]*SimpleObject)(unsafe.Pointer(GetSlice()))

	/*
		for i := 0; i < 5000; i++ {
			*realSlice = append(*realSlice, &SimpleObject{Index: int64(i + 1)})
		}
	*/
	//fmt.Printf("in-process: len: %d cap: %d\n", len(commonPool), cap(commonPool))
	ResetSlice()
	//fmt.Printf("after: len: %d cap: %d\n", len(commonPool), cap(commonPool))
}

func doWithNormal() {
	_ = make([]*SimpleObject, 0, 8192)
	/*
		for i := 0; i < 5000; i++ {
			slice = append(slice, &SimpleObject{Index: int64(i + 1)})
		}
	*/
}

func TestSliceReuse(t *testing.T) {
	doWithPool()
	doWithPool()
}

func BenchmarkTestSliceReuse(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		doWithPool()
	}
}

func BenchmarkTestSliceNormal(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		doWithNormal()
	}
}
