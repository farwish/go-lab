package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	/**
	// 应用示例，
	// 1、取消信号：
	// 当你使用 context.WithCancel(parent Context) 创建一个可取消的上下文时，它会返回两个值：一个新创建的 Context 和一个 CancelFunc。
	// CancelFunc 是一个函数，调用它可以取消这个上下文，同时也会影响到所有从这个上下文派生出的子上下文。
	// 所以下面的 cancel() 是一个函数，它是 context.CancelFunc 类型的一个实例。
	// 在这个例子中，cancel() 被用来通知 worker goroutine 上下文已经被取消，这样它就可以优雅地退出。
	*/
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	// 模拟一段时间后取消
	time.Sleep(2 * time.Second)
	cancel() // 取消上下文

	// 等待worker goroutine退出
	time.Sleep(1 * time.Second)
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker received cancellation signal.")
			return
		default:
			// 执行工作
			fmt.Println("Working...")
			time.Sleep(100 * time.Millisecond)
		}
	}
}