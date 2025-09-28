package main

import (
	"fmt"
	"sync"
	"time"
)

// Fan-Out: توزیع کار بین چندین worker
func fanOut(input <-chan int, outputs []chan int) {
	var wg sync.WaitGroup

	for data := range input {
		wg.Add(1)
		go func(d int) {
			defer wg.Done()

			// توزیع روی تمام channels (یا با الگوریتم خاص)
			for _, out := range outputs {
				out <- d
			}
		}(data)
	}

	wg.Wait()
	for _, out := range outputs {
		close(out)
	}
}

// Fan-In: جمع‌آوری نتایج از چندین منبع
func fanIn(inputs []<-chan int) <-chan int {
	output := make(chan int)
	var wg sync.WaitGroup

	for _, input := range inputs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for data := range ch {
				output <- data
			}
		}(input)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

func processor(id int, input <-chan int, output chan<- int) {
	for data := range input {
		// پردازش داده
		result := data * 2
		fmt.Printf("Processor %d: %d -> %d\n", id, data, result)
		output <- result
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// تنظیمات
	const numProcessors = 3
	const numItems = 10

	// کانال‌ها
	input := make(chan int)
	processorInputs := make([]chan int, numProcessors)
	processorOutputs := make([]<-chan int, numProcessors)

	for i := 0; i < numProcessors; i++ {
		processorInputs[i] = make(chan int)
		processorOutputs[i] = make(<-chan int)
	}

	// راه‌اندازی Fan-Out
	go fanOut(input, processorInputs)

	// راه‌اندازی processors
	for i := 0; i < numProcessors; i++ {
		outputCh := make(chan int)
		processorOutputs[i] = outputCh
		go processor(i+1, processorInputs[i], outputCh)
	}

	// راه‌اندازی Fan-In
	finalOutput := fanIn(processorOutputs)

	// ارسال داده
	go func() {
		for i := 1; i <= numItems; i++ {
			input <- i
		}
		close(input)
	}()

	// دریافت نتایج نهایی
	for result := range finalOutput {
		fmt.Printf("نتیجه نهایی: %d\n", result)
	}
}
