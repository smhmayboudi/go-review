package main

import (
	"fmt"
	"time"
)

func main() {
	// اجرای عادی (همگام)
	normalFunction()

	// اجرای Goroutine (ناهمگام)
	go goroutineFunction()

	time.Sleep(1 * time.Second)
	fmt.Println("تابع main ادامه می‌یابد")
}

func normalFunction() {
	fmt.Println("تابع عادی اجرا شد")
}

func goroutineFunction() {
	fmt.Println("Goroutine اجرا شد")
}
