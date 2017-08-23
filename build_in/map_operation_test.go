package build_in

import "testing"

var size = 50000

func assignMap() {
	m := make(map[int]int)
	for i := 0; i < size; i++ {
		m[i] = i
	}
}

func assignMapWithCap0() {
	m := make(map[int]int, size*2)
	for i := 0; i < size; i++ {
		m[i] = i
	}
}

func assignMapWithCap() {
	m := make(map[int]int, size)
	for i := 0; i < size; i++ {
		m[i] = i
	}
}

func assignMapWithCap2() {
	m := make(map[int]int, size)
	for i := 0; i < size*2; i++ {
		m[i] = i
	}
}

func BenchmarkAssignMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		assignMap()
	}
}

func BenchmarkAssignMapWithCap0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		assignMapWithCap0()
	}
}

func BenchmarkAssignMapWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		assignMapWithCap()
	}
}

func BenchmarkAssignMapWithCap2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		assignMapWithCap2()
	}
}
