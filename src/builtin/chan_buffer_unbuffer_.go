package main

func main() {
	ch := make(chan int) // 非缓冲通道

	// 非缓冲通道示例：
	go func() {
		ch <- 1 // 此处会阻塞，直到另一个goroutine执行<-ch
	}()

	ch1 := make(chan int, 3) // 缓冲通道，容量为3

	// 缓冲通道示例：
	for i := 0; i < 3; i++ {
		ch1 <- i // 这三个发送操作不会阻塞，因为缓冲区还没满
	}
	ch1 <- 4 // 此处会阻塞，因为缓冲区已满，需要等待接收
}