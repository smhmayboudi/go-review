package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// طراحی خوب: interface برای تست‌پذیری
type Processor interface {
	Process(ctx context.Context, data interface{}) (interface{}, error)
}

// طراحی خوب: ساختارهای کوچک و متمرکز
type DataProcessor struct {
	timeout time.Duration
}

func (dp *DataProcessor) Process(ctx context.Context, data interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, dp.timeout)
	defer cancel()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(100 * time.Millisecond): // شبیه‌سازی پردازش
		return fmt.Sprintf("پردازش شده: %v", data), nil
	}
}

// طراحی خوب: استفاده از sync.Pool برای کاهش تخصیص حافظه
var processorPool = sync.Pool{
	New: func() interface{} {
		return &DataProcessor{timeout: time.Second}
	},
}

func main() {
	// استفاده از pool
	processor := processorPool.Get().(*DataProcessor)
	defer processorPool.Put(processor)

	ctx := context.Background()
	result, err := processor.Process(ctx, "test data")
	if err != nil {
		fmt.Printf("خطا: %v\n", err)
	} else {
		fmt.Printf("نتیجه: %v\n", result)
	}
}
