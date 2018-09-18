package simple_impl

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

var l1 []*AA
var l2 []AA

var size = 1024 * 128

func init() {
	var size = 1024 * 1024
	l1 = make([]*AA, size)
	l2 = make([]AA, size)

	for i := 0; i < size; i++ {
		a := AA{
			v: rand.Int(),
		}
		l1[i] = &a
		l2[i] = a
	}
}

func TestPassV(t *testing.T) {
	a := AA{
		v: 200,
	}
	pass(a, 100)

	assert.Equal(t, 200, a.v)
}

func BenchmarkM(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m()
	}
}

func BenchmarkDoCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doCount(l1)
	}
}

func BenchmarkDoCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doCount2(l2)
	}
}


func BenchmarkMSlice1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mSlice1(size)
	}
}

func BenchmarkMSlice2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mSlice2(size)
	}
}

