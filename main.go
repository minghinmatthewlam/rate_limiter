package main

import (
	"github.com/minghinmatthewlam/rate_limiter/ratelimiter"
	"time"
)

func main() {
	r := ratelimiter.NewRateLimiter(3, 3*time.Second)
	r.Start()
	time.Sleep(7 * time.Second)
	r.Call()
	r.Call()
	r.Call()
	r.Call()
	r.Stop()
}
