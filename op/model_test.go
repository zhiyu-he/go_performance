package op

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

var m = &BidModel{
	field0:  200,
	field10: 100,
}

func TestGetValue(t *testing.T) {
	assert.Equal(t, int64(200), m.GetValue(unsafe.Offsetof(m.field0)))
	assert.Equal(t, int64(100), m.GetValue(unsafe.Offsetof(m.field10)))
}

func BenchmarkGetValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.GetValue(unsafe.Offsetof(m.field0))
	}
}
