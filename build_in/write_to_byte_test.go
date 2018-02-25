package build_in

import (
	"encoding/binary"
	"testing"
	"math/rand"
)


var num int64 = rand.Int63()

func buildInWrite(buf []byte, num int64) {
	binary.BigEndian.PutUint64(buf, uint64(num))
}

func buildInWrite2(b []byte, v int64) {
	_ = b[7] // early bounds check to guarantee safety of writes below
	b[0] = byte(v >> 56)
	b[1] = byte(v >> 48)
	b[2] = byte(v >> 40)
	b[3] = byte(v >> 32)
	b[4] = byte(v >> 24)
	b[5] = byte(v >> 16)
	b[6] = byte(v >> 8)
	b[7] = byte(v)
}

func selfWrite(buf []byte, num int64) {
	copy(buf, []byte{
				byte(num >> 56),
				byte(num >> 48),
				byte(num >> 40),
				byte(num >> 32),
				byte(num >> 24),
				byte(num >> 16),
				byte(num >> 8),
				byte(num)})
}


func BenchmarkWriteI64A(b *testing.B) {
	buf := make([]byte, 8)
	for i := 0; i < b.N; i++ {
		buildInWrite(buf, num)
	}
}

func BenchmarkWriteI64A2(b *testing.B) {
	buf := make([]byte, 8)
	for i := 0; i < b.N; i++ {
		buildInWrite2(buf, num)
	}
}

func BenchmarkWriteI64B(b *testing.B) {
	buf := make([]byte, 8)
	for i := 0; i < b.N; i++ {
		selfWrite(buf, num)
	}
}
