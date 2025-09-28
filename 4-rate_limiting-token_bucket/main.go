package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	tokens chan struct{}
	stop   chan struct{}
}

func NewRateLimiter(rate int, burst int) *RateLimiter {
	rl := &RateLimiter{
		tokens: make(chan struct{}, burst),
		stop:   make(chan struct{}),
	}

	// پر کردن اولیه
	for i := 0; i < burst; i++ {
		rl.tokens <- struct{}{}
	}

	// تولید token با نرخ مشخص
	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(rate))
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				select {
				case rl.tokens <- struct{}{}:
				default: // bucket پر است
				}
			case <-rl.stop:
				return
			}
		}
	}()

	return rl
}

func (rl *RateLimiter) Allow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

func (rl *RateLimiter) Wait() {
	<-rl.tokens
}

func (rl *RateLimiter) Stop() {
	close(rl.stop)
}

func main() {
	// محدود کردن به 2 درخواست در ثانیه، با burst تا 5
	limiter := NewRateLimiter(2, 5)
	defer limiter.Stop()

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func(requestID int) {
			defer wg.Done()

			// منتظر مجوز
			limiter.Wait()

			fmt.Printf("Request %d اجرا شد (زمان: %v)\n",
				requestID, time.Now().Format("15:04:05"))

			// شبیه‌سازی کار
			time.Sleep(500 * time.Millisecond)
		}(i)

		time.Sleep(200 * time.Millisecond) // فاصله بین ایجاد درخواست‌ها
	}

	wg.Wait()
}
