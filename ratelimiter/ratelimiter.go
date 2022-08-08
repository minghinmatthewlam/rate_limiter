package ratelimiter

import (
	"errors"
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

func (r *RateLimiter) Call() error {
	r.RLock()
	if !r.isRunning {
		return errors.New("rate limiter is not running")
	}

	if r.count == r.limit {
		fmt.Println("Rate limited with limit:", r.limit)
		return errors.New("rate limited")
	}
	r.RUnlock()
	r.Lock()
	defer r.Unlock()
	r.count += 1
	fmt.Println("Incremented rate limit count to", r.count)
	return nil
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

func NewRateLimiter(limit int, interval time.Duration) (*RateLimiter, error) {
	if interval == 0*time.Second {
		return nil, errors.New("cannot have zero interval rate limiter")
	}
	return &RateLimiter{
		limit:    limit,
		interval: interval,
	}, nil
}
