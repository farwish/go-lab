package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 1.WaitGroup用于等待goroutine的集合完成
	var wg sync.WaitGroup

	// 2.设置数量2个goroutine需要等待
	wg.Add(2)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("1号完成")
		// 3.goroutine执行结束调用Done
		wg.Done()
	}()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2号完成")
		wg.Done()
	}()

	// 4.阻塞直到子goroutine都结束
	wg.Wait()

	fmt.Println("都完成了")
}
