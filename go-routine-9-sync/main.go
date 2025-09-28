package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (sc *SafeCounter) Increment() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.value++
}

func (sc *SafeCounter) Value() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.value
}

func main() {
	counter := SafeCounter{}

	var wg sync.WaitGroup
	const numGoroutines = 1000

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Printf("مقدار نهایی: %d\n", counter.Value())

	// استفاده از atomic برای عملکرد بهتر
	var atomicValue int64
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&atomicValue, 1)
		}()
	}

	wg.Wait()
	fmt.Printf("مقدار atomic نهایی: %d\n", atomicValue)
}
