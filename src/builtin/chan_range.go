package main

import (
	"fmt"
	"time"
)

func main() {
	type PullRequest struct {
		Model    string `json:"model"`
		Insecure bool   `json:"insecure,omitempty"`
		Username string `json:"username"`
		Password string `json:"password"`
		Stream   *bool  `json:"stream,omitempty"`
	}
	var req PullRequest
	req.Stream = new(bool)
	*req.Stream = false

	ch := make(chan any) // 创建一个无缓冲通道
	go func() {
		defer close(ch) // 关闭通道

		for i := 0; i < 3; i++ {
			ch <- i                     // 向通道发送值
			time.Sleep(1 * time.Second) // 模拟慢速任务
		}
	}()

	if req.Stream != nil && !*req.Stream {
		// waitForStream(c, ch)
		// 同步阻塞
		for v := range ch { // 对通道进行 range 操作
			fmt.Println(v)
		}
		return
	}

	// streamResponse(c, ch)
}