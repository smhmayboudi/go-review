package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const numJobs = 20
	const numWorkers = 3 // محدود کردن تعداد workerها

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// ایجاد workerها
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// ارسال jobها
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// منتظر ماندن برای اتمام workerها
	wg.Wait()
	close(results)

	// جمع‌آوری نتایج
	for result := range results {
		fmt.Printf("نتیجه نهایی: %d\n", result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d شروع job %d\n", id, job)
		time.Sleep(1 * time.Second) // شبیه‌سازی کار سنگین
		result := job * 2
		results <- result
		fmt.Printf("Worker %d اتمام job %d, نتیجه: %d\n", id, job, result)
	}
}
