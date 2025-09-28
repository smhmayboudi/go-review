package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type WorkerPool struct {
	jobs          chan Job
	results       chan Result
	maxWorkers    int
	activeWorkers int32
	wg            sync.WaitGroup
}

type Job struct {
	ID   int
	Data interface{}
}

type Result struct {
	JobID  int
	Output interface{}
	Error  error
}

func NewWorkerPool(maxWorkers, jobQueueSize int) *WorkerPool {
	return &WorkerPool{
		jobs:       make(chan Job, jobQueueSize),
		results:    make(chan Result, jobQueueSize),
		maxWorkers: maxWorkers,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.maxWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i + 1)
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	atomic.AddInt32(&wp.activeWorkers, 1)
	defer atomic.AddInt32(&wp.activeWorkers, -1)

	fmt.Printf("Worker %d شروع به کار کرد\n", id)

	for job := range wp.jobs {
		fmt.Printf("Worker %d در حال پردازش job %d\n", id, job.ID)

		// شبیه‌سازی کار پردازشی
		time.Sleep(500 * time.Millisecond)

		// پردازش job
		result := wp.processJob(job)
		wp.results <- result
	}

	fmt.Printf("Worker %d پایان کار\n", id)
}

func (wp *WorkerPool) processJob(job Job) Result {
	// شبیه‌سازی پردازش مختلف بر اساس نوع داده
	switch data := job.Data.(type) {
	case int:
		return Result{JobID: job.ID, Output: data * data}
	case string:
		return Result{JobID: job.ID, Output: fmt.Sprintf("پردازش: %s", data)}
	default:
		return Result{JobID: job.ID, Error: fmt.Errorf("نوع داده نامعتبر")}
	}
}

func (wp *WorkerPool) SubmitJob(job Job) {
	wp.jobs <- job
}

func (wp *WorkerPool) GetResults() <-chan Result {
	return wp.results
}

func (wp *WorkerPool) Stop() {
	close(wp.jobs)
	wp.wg.Wait()
	close(wp.results)
}

func (wp *WorkerPool) GetActiveWorkers() int {
	return int(atomic.LoadInt32(&wp.activeWorkers))
}

// مثال استفاده
func main() {
	// ایجاد worker pool با 4 worker و صف 100 تایی
	pool := NewWorkerPool(4, 100)
	pool.Start()

	// ارزان 20 job
	go func() {
		for i := 1; i <= 20; i++ {
			var job Job
			if i%2 == 0 {
				job = Job{ID: i, Data: i * 10}
			} else {
				job = Job{ID: i, Data: fmt.Sprintf("کار-%d", i)}
			}
			pool.SubmitJob(job)
			fmt.Printf("Job %d ارسال شد\n", i)
		}
	}()

	// خواندن نتایج
	go func() {
		for result := range pool.GetResults() {
			if result.Error != nil {
				fmt.Printf("خطا در job %d: %v\n", result.JobID, result.Error)
			} else {
				fmt.Printf("نتیجه job %d: %v\n", result.JobID, result.Output)
			}
		}
	}()

	// منتظر ماندن برای اتمام کارها
	time.Sleep(3 * time.Second)
	pool.Stop()
	fmt.Println("همه کارها تکمیل شد")
}
