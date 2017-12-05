package build_in

import "testing"

var fillData = make([]uint64, 100000, 100000)

var preAlloc = make([]uint64, 0, 100000)
var preAlloc2 = make([]uint64, 100000, 100000)

func copyWithNilArray() {
	var array []uint64
	if len(array) != len(fillData) {
		array = make([]uint64, len(fillData))
	}
	copy(array, fillData)
}


func copyWithPreAllocArray(dat []uint64) {
	copy(preAlloc2, dat)
}

func appendWithPreAllocArray(dat []uint64) {
	preAlloc = append(preAlloc, dat...)
	preAlloc = preAlloc[0:0]
}

func BenchmarkCopyWithNilArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		copyWithNilArray()
	}
}

func BenchmarkCopyWithNilArray2(b *testing.B) {
	BenchmarkCopyWithNilArray(b)
}

func BenchmarkCopyWithNilArray3(b *testing.B) {
	BenchmarkCopyWithNilArray2(b)
}

func BenchmarkAppendWithPreAllocArray(b *testing.B) {
	dat := []uint64{1}
	for i := 0; i < b.N; i++ {
		appendWithPreAllocArray(dat)
	}
}

func BenchmarkAppendWithPreAllocArray2(b *testing.B) {
	BenchmarkAppendWithPreAllocArray(b)
}

func BenchmarkAppendWithPreAllocArray3(b *testing.B) {
	BenchmarkAppendWithPreAllocArray2(b)
}

func BenchmarkCopyWithPreArray(b *testing.B) {
	dat := []uint64{1}
	for i := 0; i < b.N; i++ {
		copyWithPreAllocArray(dat)
	}
}

func BenchmarkCopyWithPreArray2(b *testing.B) {
	BenchmarkCopyWithPreArray(b)
}

func BenchmarkCopyWithPreArray3(b *testing.B) {
	BenchmarkCopyWithPreArray(b)
}
