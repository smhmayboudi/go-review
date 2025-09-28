package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID   int
	Data string
}

type Result struct {
	JobID  int
	Output string
	Error  error
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d شروع job %d\n", id, job.ID)

		// شبیه‌سازی پردازش
		time.Sleep(1 * time.Second)

		result := Result{
			JobID:  job.ID,
			Output: fmt.Sprintf("پردازش شده: %s", job.Data),
		}

		results <- result
		fmt.Printf("Worker %d پایان job %d\n", id, job.ID)
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	// راه‌اندازی workerها
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// ارسال jobs
	go func() {
		for i := 1; i <= numJobs; i++ {
			jobs <- Job{
				ID:   i,
				Data: fmt.Sprintf("داده-%d", i),
			}
		}
		close(jobs)
	}()

	// جمع‌آوری نتایج
	go func() {
		wg.Wait()
		close(results)
	}()

	// نمایش نتایج
	for result := range results {
		fmt.Printf("نتیجه: %+v\n", result)
	}
}
