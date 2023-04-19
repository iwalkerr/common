package limitx

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	burst = 1000 // 令牌桶最大值
	rate  = 1000 // 每秒生成几个令牌

	seconds = 10 // 测试时间为5秒
)

func TokenLimit() {
	store := redis.New("0.0.0.0:6379")
	fmt.Println(store.Ping())
	// New tokenLimiter
	limiter := limit.NewTokenLimiter(rate, burst, store, "tokenlimit:")

	timer := time.NewTimer(time.Second * time.Duration(seconds))
	quit := make(chan struct{})
	defer timer.Stop()
	go func() {
		<-timer.C
		close(quit)
	}()

	var allowed, denied int32
	var wait sync.WaitGroup
	for i := 0; i < runtime.NumCPU(); i++ {
		wait.Add(1)
		go func(num int) {
			for {
				select {
				case <-quit:
					wait.Done()
					return
				default:
					if limiter.Allow() {
						// logx.Infof("当前是 %d CPU执行了", num)
						atomic.AddInt32(&allowed, 1)
					} else {
						atomic.AddInt32(&denied, 1)
					}
				}
			}
		}(i)
	}

	wait.Wait()
	fmt.Printf("allowed: %d, denied: %d, qps: %d\n", allowed, denied, (allowed+denied)/seconds)
}
