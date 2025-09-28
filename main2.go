package main

import "fmt"

func main() {
	// ایجاد یک Channel برای传递 مقادیر int
	ch := make(chan int)

	// یک Goroutine برای ارسال داده
	go func() {
		result := 100 * 5
		ch <- result // ارسال نتیجه به Channel
	}()

	// دریافت داده از Channel (این خط منتظر می‌ماند تا داده ای برسد)
	value := <-ch
	fmt.Println(value) // 500
}
