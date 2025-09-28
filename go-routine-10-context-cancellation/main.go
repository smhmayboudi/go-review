package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int, results chan<- int) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d متوقف شد: %v\n", id, ctx.Err())
			return
		case results <- i:
			fmt.Printf("Worker %d تولید کرد: %d\n", id, i)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	results := make(chan int)

	// راه‌اندازی workerها
	for i := 1; i <= 3; i++ {
		go worker(ctx, i, results)
	}

	// خواندن نتایج تا timeout
	for {
		select {
		case <-ctx.Done():
			fmt.Println("برنامه به پایان رسید")
			return
		case result := <-results:
			fmt.Printf("نتیجه دریافت شد: %d\n", result)
		}
	}
}
