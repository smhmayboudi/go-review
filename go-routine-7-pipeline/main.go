package main

import (
	"fmt"
	"time"
)

// Stage 1: تولید اعداد
func generate(numbers ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numbers {
			out <- n
		}
		close(out)
	}()
	return out
}

// Stage 2: مربع اعداد
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// Stage 3: چاپ نتایج
func print(in <-chan int) {
	for result := range in {
		fmt.Printf("نتیجه: %d\n", result)
		time.Sleep(100 * time.Millisecond)
	}
}

// Pipeline کامل
func main() {
	// ساختار: generate -> square -> print
	numbers := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	squared := square(numbers)
	print(squared)
}
