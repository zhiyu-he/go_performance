package build_in

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var (
	rawData1 = []int64{1, 2, 3}
	rawData2 = []int64{4, 5, 6}
	rawData3 = []int64{7, 8, 9}

	empty1 = []int64{}
	empty2 = []int64{}
	empty3 = []int64{}

)


func TestEmptyAlloc(t *testing.T) {
	l := appendEmpty()
	assert.Equal(t, 0, l)
}

func appendEmpty() int {
	tmp := append(empty1, append(empty2, empty3...)...)
	return len(tmp)
}

func appendTwoArray() int {
	m := append([]int64{1,2,3}, []int64{4, 5, 6}...)
	q := append(m, append([]int64{1,2,3}, []int64{4, 5, 6}...)...)
	return len(q)
}

func appendMethod1() {
	cnt := 0
	for _, num := range append(rawData1, append(rawData2, rawData3...)...) {
		if num > 0 {
			cnt += 1
		}
	}
}

func appendMethod2() {
	tmp := make([]int64, 0, len(rawData1)+len(rawData2)+len(rawData3))
	tmp = append(tmp, rawData1...)
	tmp = append(tmp, rawData2...)
	tmp = append(tmp, rawData3...)

	cnt := 0
	for _, num := range tmp {
		if num > 0 {
			cnt += 1
		}
	}
}

func TestBase(t *testing.T) {
	tmp := make([]int64, 0, len(rawData1)+len(rawData2)+len(rawData3))
	tmp = append(tmp, rawData1...)
	tmp = append(tmp, rawData2...)
	tmp = append(tmp, rawData3...)

	assert.Equal(t, 9, len(tmp))
}

func TestAppendTwoArray(t *testing.T) {
	length := appendTwoArray()
	assert.Equal(t, 12, length)
}

func BenchmarkAppendTwoArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appendTwoArray()
	}
}

func BenchmarkAppendMethod1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appendMethod1()
	}
}

func BenchmarkAppendMethod2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appendMethod2()
	}
}

func BenchmarkAppendEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		appendEmpty()
	}
}