package ratelimiter

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestZeroInterval(t *testing.T) {
	_, err := NewRateLimiter(0, 0*time.Second)
	assert.Error(t, err)
}

func TestNotStarted(t *testing.T) {
	limit := 3
	r, err := NewRateLimiter(limit, 3*time.Second)
	assert.NoError(t, err)
	assert.Error(t, r.Call())
}

func TestRateLimited(t *testing.T) {
	limit := 3
	r, err := NewRateLimiter(limit, 3*time.Second)
	assert.NoError(t, err)
	r.Start()
	for i := 0; i < limit; i++ {
		err = r.Call()
		assert.NoError(t, err)
	}
	err = r.Call()
	assert.Error(t, err)
}

func TestLimitReset(t *testing.T) {
	limit := 3
	interval := 3 * time.Second
	r, err := NewRateLimiter(limit, interval)
	assert.NoError(t, err)
	r.Start()
	for i := 0; i < limit; i++ {
		err = r.Call()
		assert.NoError(t, err)
	}

	time.Sleep(interval + time.Second)
	assert.NoError(t, r.Call())
}
