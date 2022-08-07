package ratelimiter

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	limit     int
	interval  time.Duration
	count     int
	isRunning bool
	sync.RWMutex
}

func (r *RateLimiter) Call() bool {
	r.RLock()
	if r.count == r.limit {
		fmt.Println("Rate limited with limit:", r.limit)
		return false
	}
	r.RUnlock()
	r.Lock()
	defer r.Unlock()
	r.count += 1
	fmt.Println("Incremented rate limit count to", r.count)
	return true
}

func (r *RateLimiter) Start() {
	r.isRunning = true
	go func() {
		for r.isRunning {
			fmt.Println("Rate limiter started")
			time.Sleep(r.interval)
			r.Reset()
		}
	}()
	fmt.Println("End of rate limit start func")
}

func (r *RateLimiter) Reset() {
	r.Lock()
	defer r.Unlock()
	r.count = 0
	fmt.Println("Reset rate limiter count")
}

func (r *RateLimiter) Stop() {
	r.isRunning = false
	fmt.Println("Rate limiter stopped")
}

func NewRateLimiter(limit int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:    limit,
		interval: interval,
	}
}
