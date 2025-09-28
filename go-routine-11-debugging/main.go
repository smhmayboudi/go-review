package main

import (
	"fmt"
	"log"
	"time"
)

func safeGoroutine(id int, panicMode bool) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Goroutine %d panic بازیابی شد: %v", id, r)
		}
	}()

	if panicMode {
		panic(fmt.Sprintf("panic عمدی در goroutine %d", id))
	}

	fmt.Printf("Goroutine %d در حال اجرا\n", id)
	time.Sleep(1 * time.Second)
}

func main() {
	// اجرای ایمن
	go safeGoroutine(1, false)

	// اجرای با panic (اما بازیابی شده)
	go safeGoroutine(2, true)

	time.Sleep(2 * time.Second)
	fmt.Println("برنامه ادامه می‌یابد")
}
