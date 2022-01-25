package main

import (
	"fmt"
)

func main() {
	/*
		Channel创建方式：
		var c1 chan [value type]
		c1 = make([channel type] [value type], [capacity])

		[value type] 定义的是 Channel 中所传输数据的类型。
		[channel type] 定义的是 Channel 的类型，其类型有以下三种：
				"chan" 	  可读可写 : chan int 则表示可读写 int 数据的 channel
				"chan<-"  仅可写   : chan<- float64 则表示仅可写64位 float 数据的 channel
				"<-chan"  仅可读   : <-chan int 则表示仅可读 int 数据的 channel
		[capacity] 是一个可选参数，其定义的是 channel 中的缓存区 (buffer)。
					  如果不填则默认该 channel 没有缓冲区 (unbuffered)。
					  对于没有缓冲区的 channel，消息的发送和收取必须能同时完成，否则会造成阻塞并提示死锁错误。
		chan读写使用的符号都是<-
	*/

	var c1 chan int
	c1 = make(chan int, 100)

	// 1.向 Channel1 发送(写入)10
	c1 <- 10

	// 2.从 Channel1 接收(读取)
	i := <-c1

	fmt.Println("Receive from channel 1: ", i)

	// ==================================================================

	/*
		Channel死锁：对 channel 的发送和接收动作永远不会同时发生，从而阻塞造成死锁。
		fatal error: all goroutines are asleep - deadlock!
	*/
	var c2 chan int
	c2 = make(chan int)

	// 没有buffer的channel，先写后读 会出现死锁：
	// func() {
	// 	c2 <- 0
	// }()
	//
	// j := <-c2

	// 避免死锁方式1：使用goroutine并发执行
	// 通过 go 语句定义发送操作的方程在另一个协程并发运行，chan读取没有数据时会阻塞等待，从而能够解决死锁问题。
	go func() {
		c2 <- 2
	}()

	j := <-c2

	fmt.Println("Receive from channel 2: ", j)

	// 避免死锁方式2：使用buffer
	// 为 channel 添加一个缓冲区（buffer），这样只要 buffer 没有用尽，阻塞就不会发生，死锁也不会发生。
	var c3 chan int
	c3 = make(chan int, 1)

	func() {
		c3 <- 3
	}()

	k := <-c3

	fmt.Println("Receive from channel 3: ", k)
}
