package main

import (
	"fmt"
	"runtime"
)

func showSchedulerInfo() {
	fmt.Printf("تعداد CPU: %d\n", runtime.NumCPU())
	fmt.Printf("تعداد Goroutines: %d\n", runtime.NumGoroutine())
	fmt.Printf("نسخه Go: %s\n", runtime.Version())
}

func main() {
	showSchedulerInfo()

	for i := 0; i < 10; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d - تعداد Goroutines: %d\n",
				id, runtime.NumGoroutine())
		}(i)
	}

	runtime.Gosched() // اجازه اجرا به Goroutines دیگر
	showSchedulerInfo()
}
