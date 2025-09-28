package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type MonitoredWorkerPool struct {
	jobs       chan int
	results    chan int
	maxWorkers int

	// متریک‌ها
	totalJobs     int64
	processedJobs int64
	activeWorkers int32
	errors        int64
}

func NewMonitoredWorkerPool(maxWorkers int) *MonitoredWorkerPool {
	return &MonitoredWorkerPool{
		jobs:       make(chan int, 100),
		results:    make(chan int, 100),
		maxWorkers: maxWorkers,
	}
}

func (mwp *MonitoredWorkerPool) Start() {
	for i := 0; i < mwp.maxWorkers; i++ {
		go mwp.worker(i)
	}

	// شروع monitoring
	go mwp.monitor()
}

func (mwp *MonitoredWorkerPool) worker(id int) {
	atomic.AddInt32(&mwp.activeWorkers, 1)
	defer atomic.AddInt32(&mwp.activeWorkers, -1)

	for job := range mwp.jobs {
		// پردازش job
		result, err := mwp.processJob(job)
		if err != nil {
			atomic.AddInt64(&mwp.errors, 1)
			continue
		}

		mwp.results <- result
		atomic.AddInt64(&mwp.processedJobs, 1)
	}
}

func (mwp *MonitoredWorkerPool) processJob(job int) (int, error) {
	if job < 0 {
		return 0, fmt.Errorf("job منفی مجاز نیست")
	}
	time.Sleep(100 * time.Millisecond)
	return job * 2, nil
}

func (mwp *MonitoredWorkerPool) SubmitJob(job int) {
	mwp.jobs <- job
	atomic.AddInt64(&mwp.totalJobs, 1)
}

func (mwp *MonitoredWorkerPool) monitor() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		total := atomic.LoadInt64(&mwp.totalJobs)
		processed := atomic.LoadInt64(&mwp.processedJobs)
		active := atomic.LoadInt32(&mwp.activeWorkers)
		errors := atomic.LoadInt64(&mwp.errors)

		fmt.Printf("[Monitoring] Active: %d, Total: %d, Processed: %d, Errors: %d, Queue: %d\n",
			active, total, processed, errors, len(mwp.jobs))
	}
}

func main() {
	pool := NewMonitoredWorkerPool(3)
	pool.Start()

	// ارسال کارها
	for i := 1; i <= 20; i++ {
		pool.SubmitJob(i)
		time.Sleep(50 * time.Millisecond)
	}

	time.Sleep(3 * time.Second)
}
