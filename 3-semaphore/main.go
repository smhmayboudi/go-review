package main

import (
	"fmt"
	"sync"
	"time"
)

// Semaphore برای محدود کردن همزمانی
type Semaphore chan struct{}

func NewSemaphore(n int) Semaphore {
	return make(chan struct{}, n)
}

func (s Semaphore) Acquire() {
	s <- struct{}{}
}

func (s Semaphore) Release() {
	<-s
}

func (s Semaphore) TryAcquire() bool {
	select {
	case s <- struct{}{}:
		return true
	default:
		return false
	}
}

func main() {
	const maxConcurrent = 3
	const totalTasks = 10

	sem := NewSemaphore(maxConcurrent)
	var wg sync.WaitGroup

	for i := 1; i <= totalTasks; i++ {
		wg.Add(1)

		go func(taskID int) {
			defer wg.Done()

			// کسب مجوز
			sem.Acquire()
			defer sem.Release()

			fmt.Printf("Task %d شروع شد (همزمانی فعال: %d/%d)\n",
				taskID, len(sem), maxConcurrent)

			// شبیه‌سازی کار
			time.Sleep(2 * time.Second)

			fmt.Printf("Task %d پایان یافت\n", taskID)
		}(i)
	}

	wg.Wait()
	fmt.Println("همه tasks تکمیل شد")
}
