package main

import (
	"encoding/binary"
	"math/rand"
	"fmt"
)

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


func main() {
	b := make([]byte, 1024)
	num := rand.Int63()
	idx := 0
	buildInWrite(b, num)
	idx += 8
	selfWrite(b[idx:], num)
	idx += 8
	buildInWrite2(b[idx:], num)
	idx += 8

	fmt.Printf("%v\n", b[:idx])
}
