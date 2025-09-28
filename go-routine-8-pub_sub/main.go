package main

import (
	"fmt"
	"sync"
	"time"
)

type PubSub struct {
	subscribers map[string][]chan string
	mu          sync.RWMutex
}

func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: make(map[string][]chan string),
	}
}

func (ps *PubSub) Subscribe(topic string) <-chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan string, 1)
	ps.subscribers[topic] = append(ps.subscribers[topic], ch)

	return ch
}

func (ps *PubSub) Publish(topic string, message string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	if subscribers, exists := ps.subscribers[topic]; exists {
		for _, ch := range subscribers {
			go func(c chan string) {
				c <- message
			}(ch)
		}
	}
}

func main() {
	ps := NewPubSub()

	// مشترک‌ها
	sub1 := ps.Subscribe("news")
	sub2 := ps.Subscribe("sports")
	sub3 := ps.Subscribe("news") // مشترک دوم برای news

	// انتشار پیام
	go func() {
		ps.Publish("news", "خبر فوری: Go 1.21 منتشر شد!")
		ps.Publish("sports", "نتایج مسابقات فوتبال")
		ps.Publish("news", "اخبار اقتصادی روز")
	}()

	// دریافت پیام‌ها
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for msg := range sub1 {
			fmt.Printf("مشترک 1 (news): %s\n", msg)
		}
	}()

	go func() {
		defer wg.Done()
		for msg := range sub2 {
			fmt.Printf("مشترک 2 (sports): %s\n", msg)
		}
	}()

	go func() {
		defer wg.Done()
		for msg := range sub3 {
			fmt.Printf("مشترک 3 (news): %s\n", msg)
		}
	}()

	time.Sleep(1 * time.Second)
	wg.Done()
}
