package simple_impl

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"time"
)

func TestCreate(t *testing.T) {
	limiter := Create(5)
	for i := 0; i < 5; i++ {
		assert.Equal(t, true, limiter.TryAcquire())
	}
	assert.Equal(t, true, limiter.TryAcquire())
	assert.Equal(t, false, limiter.TryAcquire())
	assert.Equal(t, false, limiter.TryAcquire())
	time.Sleep(time.Millisecond * 200)
	t.Logf("%s\n", limiter.Dump())
	assert.Equal(t, true, limiter.TryAcquire())
	assert.Equal(t, false, limiter.TryAcquire())
}
