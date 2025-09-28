package main

import (
	"errors"
	"fmt"
	"time"
)

type Result struct {
	Value int
	Error error
}

func processData(data int) Result {
	if data < 0 {
		return Result{Error: errors.New("عدد منفی مجاز نیست")}
	}

	// شبیه‌سازی پردازش
	time.Sleep(100 * time.Millisecond)
	return Result{Value: data * 2}
}

func worker(data int, results chan<- Result) {
	result := processData(data)
	results <- result
}

func main() {
	data := []int{1, 2, -3, 4, 5, -6}
	results := make(chan Result, len(data))

	// راه‌اندازی workerها
	for _, d := range data {
		go worker(d, results)
	}

	// جمع‌آوری نتایج
	for i := 0; i < len(data); i++ {
		result := <-results
		if result.Error != nil {
			fmt.Printf("خطا: %v\n", result.Error)
		} else {
			fmt.Printf("موفق: %d\n", result.Value)
		}
	}
}
