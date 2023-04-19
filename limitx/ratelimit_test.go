package limitx

import (
	"fmt"
	"testing"
	"time"

	irate "golang.org/x/time/rate"
)

// 令牌桶
func TestRatelimit(t *testing.T) {
	// 每秒填充7个令牌
	funcRun := RateLimitMiddleware(1, 7)

	for i := 0; i < 10; i++ {
		funcRun()
	}
}

func RateLimitMiddleware(rate float64, capacity int) func() {
	limit := irate.NewLimiter(irate.Every(time.Duration(rate)*time.Second), capacity)
	return func() {
		if !limit.Allow() {
			fmt.Println("rate limit...")
			return
		}
		fmt.Println("next run")
	}
}
