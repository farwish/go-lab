package main

import "fmt"

func main() {
	ch := make(chan string) // 非缓冲通道
	// 在一个goroutine中发送数据：
	go func() {
		ch <- "hello" // 此处会阻塞，直到另一个goroutine执行<-ch
	}()
	// 主goroutine接收数据
	value := <-ch // 这里会阻塞，直到有值被发送到ch
	fmt.Println("Received from unbuffered channel:", value)


	ch1 := make(chan int, 3) // 缓冲通道，容量为3
	// 假设已经向ch1中发送了若干个整数
	for i := 0; i < cap(ch1); i++ {
		ch1 <- i // 不会阻塞 直到换冲区已满; 假设这些发送操作已经执行过
	}
	// 消费通道中的所有数据
	for i := 0; i < cap(ch1); i++ {
		value := <-ch1
		fmt.Println("Received from buffered channel:", value)
	}
}