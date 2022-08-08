package main

import (
	"fmt"
	"github.com/minghinmatthewlam/rate_limiter/ratelimiter"
	"time"
)

func main() {
	buffer := make(chan int, 100)
	for i := 0; i < 10; i++ {
		buffer <- 1
	}
	fmt.Println(len(buffer), cap(buffer))
	r, err := ratelimiter.NewRateLimiter(3, 3*time.Second)
	if err != nil {
		panic(fmt.Sprintf("Rate limiter could not be initalized with error %e", err))
	}
	r.Start()
	time.Sleep(1 * time.Second)
	r.Call()
	r.Call()
	r.Call()
	r.Call()
	r.Stop()
}
