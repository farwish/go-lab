package main

import (
	"context"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	// 创建一个带有超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// 使用 ctx 创建 errgroup
	g, ctx := errgroup.WithContext(ctx)

	// 向 errgroup 添加任务
	g.Go(func() error {
		return doSomething(ctx, "Task 1")
	})
	g.Go(func() error {
		return doSomething(ctx, "Task 2")
	})

	// 等待所有任务完成，如果有任务被取消或超时，Wait 会返回相应的错误
	if err := g.Wait(); err != nil {
		log.Printf("An error occurred: %v\n", err)
	} else {
		log.Println("All tasks completed successfully.")
	}
}

func doSomething(ctx context.Context, taskName string) error {
	// 当log.Println("/start/")和time.Sleep(3 * time.Second)放在select语句之上时，
	// 无论select中的条件ctx.Done()是否满足，这部分代码都会执行。
	// 因此，/start/会被打印，接着程序会睡眠3秒。

	// 但是，这并不意味着select语句失去了作用。
	// select中的case <-ctx.Done():分支会在time.Sleep结束后立即执行，
	// 因为time.Sleep不会阻止Go的运行时检查通道事件。
	// 这意味着即便/end/被打印，紧随其后的return ctx.Err();依然能够反映任务因上下文取消而提前终止的情况。

	// 在实际执行中，在sleep之后直接输出/end/可能在上下文取消之后仍被打印，给人一种任务成功结束的错觉。
	// 正确的逻辑处理应当依赖于函数返回的错误，而非仅凭日志输出判断。
	log.Println("/start/")
	doing()
	log.Println("/end/")

	// 使用 ctx 来检查是否应该取消操作
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		// 下面代码如果放在 default 内，和放在 select 上面的含义不一样：
		// 进入 select 语句，如果上下文没有被取消，default分支会被执行，先打印/start/，然后time.Sleep，最后打印/end/。
		// log.Println("/start/")
		// doing()
		// log.Println("/end/")

		log.Printf("%s completed.", taskName)
		return nil
	}

}

func doing() {
	time.Sleep(3 * time.Second)
}