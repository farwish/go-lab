package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	/**
	// 应用示例，
	// 2、截止时间：
	// 这个例子展示了如何设置一个超时，如果操作在指定时间内未完成，context 会自动取消。
	// 我们创建了一个新的上下文ctx，其中包含了一个名为completionKey的键和completion通道的值。
	// 在longRunningOperation中，我们从上下文中获取这个通道，并在操作完成后关闭它。这样，main函数可以通过监听completion通道来得知操作是否成功完成。
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	completionKey := "completion"
	completion := make(chan bool)
	ctx = context.WithValue(ctx, completionKey, completion)

	go longRunningOperation(ctx)

	// 等待操作完成或超时
	select {
	case <-ctx.Done():
		fmt.Println("Operation timed out or was canceled.")
	case <-completion: // 从上下文中获取完成信号
		fmt.Println("Operation completed successfully.")
	}
}

func longRunningOperation(ctx context.Context) {
	completion := ctx.Value("completion").(chan bool) // 从上下文中获取完成信号通道

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Operation canceled by context.")
			return
		default:
			// 模拟一个长运行的操作
			time.Sleep(10 * time.Second)
			fmt.Println("Operation would have completed, but checking context first.")
			close(completion) // 操作模拟完毕，关闭完成信号通道
			return // 操作模拟完毕，正常返回
		}
	}
}