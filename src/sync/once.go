package main

import (
	"fmt"
	"log"
	"sync"

	"golang.org/x/sync/errgroup"
)

func main() {
	/**
		$GOROOT/src/sync/once.go
		定义了Go语言 sync 包中的 Once 类型和其方法 Do，用于确保某个函数（初始化函数）在程序中只执行一次。

		* “Once类型” (type Once struct{xxx})包含一个done原子类型和一个m互斥锁。done用于标记初始化函数是否已经执行过。
		* “Do方法” 接受一个无参数的函数 f 作为输入，如果 Once 的 done 值为0（即初始化函数尚未执行），则调用 doSlow 方法执行函数 f。
		* “doSlow方法” 首先加锁，确保同时只有一个协程能执行初始化函数。然后检查 done 值，如果为0，则执行函数 f，并在 defer 语句中设置 done 为1，以标记初始化函数已完成。最后解锁。
		
		这个实现确保了即使在多协程环境下，初始化函数也只会被执行一次，并且能够处理并发调用 Do 方法的情况。
	*/	
	
	// 使用示例：

	var once sync.Once
	var g errgroup.Group

	initFunc := func() {
		// 定义一个需要初始化的操作
		fmt.Println("初始化完成，这个信息只打印一次")
	}

	for i := 0; i < 10; i++ {
		// 在多个goroutine中安全地调用初始化函数
		g.Go(func () error {
			log.Printf("goroutine %d", i)
			once.Do(initFunc)
			return nil
		})
	}
	if err := g.Wait(); err == nil {
        fmt.Println("Successfully done.")
    }
}