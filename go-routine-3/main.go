package main

func main() {
	// ۱. Unbuffered Channel (همگام)
	ch1 := make(chan int)

	// ۲. Buffered Channel (ناهمگام)
	ch2 := make(chan int, 3)

	// ۳. Channel فقط برای ارسال
	var sendOnly chan<- int = ch1

	// ۴. Channel فقط برای دریافت
	var receiveOnly <-chan int = ch1

	_ = ch2
	_ = sendOnly
	_ = receiveOnly
}
