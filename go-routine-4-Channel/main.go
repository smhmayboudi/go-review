package main

import (
	"fmt"
	"time"
)

// الگوی Producer-Consumer
func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("تولید: %d\n", i)
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func consumer(ch <-chan int, done chan<- bool) {
	for value := range ch {
		fmt.Printf("مصرف: %d\n", value)
		time.Sleep(200 * time.Millisecond)
	}
	done <- true
}

func main() {
	dataCh := make(chan int, 2)
	doneCh := make(chan bool)

	go producer(dataCh)
	go consumer(dataCh, doneCh)

	<-doneCh
	fmt.Println("پایان")
}
