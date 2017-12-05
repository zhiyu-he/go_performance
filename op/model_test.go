package op

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
	"fmt"
)

var m = &BidModel{
	field0:  200,
	field10: 100,
}

func TestGetValue(t *testing.T) {
	assert.Equal(t, int64(200), m.GetValue(unsafe.Offsetof(m.field0)))
	assert.Equal(t, int64(100), m.GetValue(unsafe.Offsetof(m.field10)))

	m.SetValue(unsafe.Offsetof(m.field11), 999)

	assert.Equal(t, int64(999), m.GetValue(unsafe.Offsetof(m.field11)))

	m1 := &M{}


	m.SetStructValue(unsafe.Offsetof(m.m1), m1)

	m.GetStructValue(unsafe.Offsetof(m.m1)).num = 100
	fmt.Printf("%#v\n", m.GetStructValue(unsafe.Offsetof(m.m1)))
}

func BenchmarkGetValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m.GetValue(unsafe.Offsetof(m.field0))
	}
}

