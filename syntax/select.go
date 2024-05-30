package main

import "fmt"

func main() {
	var chan1 chan string
	var chan2 chan string

	chan1 = make(chan string, 1)
	chan2 = make(chan string, 1)

	/*
		select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。 每个 case 必须是一个通信操作，要么是发送要么是接收。
		select 随机执行一个可运行的 case。 如果没有 case 可运行，它将阻塞，直到有 case 可运行。
		一个默认的子句应该总是可运行的。

		备注：可以简单理解select在做多路复用，进行通信检测切换。
	*/

	// 都阻塞，有default分支 则走default
	select {
	case <-chan1:
		fmt.Println("chan1 read success")
	case <-chan2:
		fmt.Println("chan2 write success")
	default:
		fmt.Println("default")
	}

	fmt.Println("1.---------------------")

	// 有没阻塞的，则执行该分支执行（chan2 写入）
	select {
	case <-chan1:
		fmt.Println("chan1 read success")
	case chan2 <- "1":
		fmt.Println("chan2 write success")
	}

	fmt.Println("2.---------------------")

	// 有没阻塞的，则执行该分支执行（chan2 读取）
	select {
	case <-chan1:
		fmt.Println("chan1 read success")
	case value := <-chan2:
		fmt.Println("chan2 read success", value)
	}

	fmt.Println("3.---------------------")

	// 都阻塞，且没有default分支 则死锁
	select {
	case <-chan1:
		fmt.Println("chan1 read success")
	case <-chan2:
		fmt.Println("chan2 read success")
	}
}
