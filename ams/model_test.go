package main

import (
	"math/rand"
	"testing"
)

var s []*MiniModel

func init() {
	s := make([]*MiniModel, 0, 30000)

	for i := 0; i < 30000; i++ {
		m := &MiniModel{
			ValType: int32(rand.Intn(10)),
		}
		m.Set()
		s = append(s, m)
	}
}

func BenchmarkTestIsType(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cnt := 0
		for _, m := range s {
			if m.IsXXType() {
				cnt += 1
			}
		}
	}
}
func BenchmarkTestIsType2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cnt := 0
		for _, m := range s {
			if m.IsXXTypeBool {
				cnt += 1
			}
		}
	}
}
