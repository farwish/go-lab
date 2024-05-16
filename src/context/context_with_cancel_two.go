package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	/**
	// 应用示例，
	// 4、级联取消：
	// 这个例子显示了当父上下文取消时，所有从它派生的子上下文也会被取消。
	// 这些例子展示了 context 如何在并发环境中提供取消、截止时间、数据传递和级联取消等关键功能。
	 */

	// 创建根上下文并初始化取消函数
	ctx, cancel := context.WithCancel(context.Background())
	// 从根上下文派生第一个子上下文
	childCtx, _ := context.WithCancel(ctx)

	// 启动两个工作协程，其中worker1直接使用了子上下文，worker2隐式地使用了根上下文
	go worker1(childCtx)
	go worker2(ctx)

	// 取消根上下文，这将导致所有衍生的子上下文也被取消
	cancel()

	// 主函数阻塞，等待工作协程结束，观察输出。
	time.Sleep(3 * time.Millisecond)
}

func worker1(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("Worker1 received cancellation signal.")
}

func worker2(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("Worker2 received cancellation signal.")
}