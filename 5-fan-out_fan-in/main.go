package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const numProducers = 5
	const numConsumers = 3
	const bufferSize = 10

	// کانال‌ها
	dataChan := make(chan int, bufferSize)
	resultChan := make(chan string, bufferSize)

	var producerWg sync.WaitGroup
	var consumerWg sync.WaitGroup

	// Producers (Fan-Out)
	for i := 0; i < numProducers; i++ {
		producerWg.Add(1)
		go func(id int) {
			defer producerWg.Done()
			for j := 0; j < 10; j++ {
				data := id*100 + j
				dataChan <- data
				fmt.Printf("Producer %d تولید کرد: %d\n", id, data)
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	// Consumers (Fan-In)
	for i := 0; i < numConsumers; i++ {
		consumerWg.Add(1)
		go func(id int) {
			defer consumerWg.Done()
			for data := range dataChan {
				// پردازش داده
				result := fmt.Sprintf("Consumer %d پردازش کرد: %d -> %d",
					id, data, data*2)
				resultChan <- result
				time.Sleep(200 * time.Millisecond)
			}
		}(i)
	}

	// بستن کانال بعد از اتمام producers
	go func() {
		producerWg.Wait()
		close(dataChan)
	}()

	// جمع‌آوری نتایج
	go func() {
		consumerWg.Wait()
		close(resultChan)
	}()

	// نمایش نتایج
	for result := range resultChan {
		fmt.Println(result)
	}

	fmt.Println("پردازش کامل شد")
}
