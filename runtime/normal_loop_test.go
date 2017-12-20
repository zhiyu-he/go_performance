package runtime

import (
	"math/rand"
	"testing"
	"encoding/hex"
)

type A struct {
	F1 int
	F2 int
	F3 int
	F4 int
	s1 string
	s2 string
}


var l1 []*A
var l2 []A


func init() {
	b := make([]byte, 10000)
	rand.Read(b)
	str := hex.EncodeToString(b)

	size := 10
	l1 = make([]*A, size)
	l2 = make([]A, size)

	for i := 0; i < size; i++ {
		a := A{
			F1: rand.Int(),
			F2: rand.Int(),
			F3: rand.Int(),
			F4: rand.Int(),
			s1: str,
			s2: str,
		}

		x := &a
		l1[i] = x
		l2[i] = a
	}
}


func normalLoop(l []*A) {
	var sum int = 0
	for _, item := range l {
		sum += item.F3
	}
}

func optLoop(l *[]A) {
	var sum int = 0
	for idx := range *l {
		sum += (*l)[idx].F3
	}
}


func BenchmarkRead1(b *testing.B) {
	var sum = 0
	for i := 0; i < b.N; i++ {
		for _, m := range l1 {
			sum += m.F3
		}
	}
}

func BenchmarkRead2(b *testing.B) {
	var sum = 0
	for i := 0; i < b.N; i++ {
		for _, m := range l2 {
			sum += m.F3
		}
	}
}

func BenchmarkRead3(b *testing.B) {
	var sum = 0
	for i := 0; i < b.N; i++ {
		for idx := range l2 {
			sum += l2[idx].F3
		}
	}
}


func BenchmarkLoop1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		normalLoop(l1)
	}
}

func BenchmarkLoop2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optLoop(&l2)
	}
}