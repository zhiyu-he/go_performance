package build_in

import (
	"encoding/binary"
	"testing"
)

func copy2dst(src *[]byte, dst *[]byte) {
	copy(*dst, *src)
}

var buffer [8]byte

func f1(val int64, dst *[]byte) {
	v := buffer[0:8]
	binary.BigEndian.PutUint64(v, uint64(val))
	copy2dst(&v, dst)
}

func f2(val int64, offset int, dst *[]byte) {
	binary.BigEndian.PutUint64((*dst)[offset:], uint64(val))
}

var dst1 []byte = make([]byte, 1024*1024*100)
var dst2 []byte = make([]byte, 1024*1024*100)

func Loop1() {
	for i := 0; i < 36000; i++ {
		f1(int64(i), &dst1)
	}
}

func Loop2() {
	idx := 0
	for i := 0; i < 36000; i++ {
		f2(int64(i), idx+i*8, &dst2)
	}
}

func BenchmarkLoop1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop1()
	}
}

func BenchmarkLoop2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop2()
	}
}
