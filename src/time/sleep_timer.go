package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个定时器，设置初始时间为2秒
	t := time.NewTimer(3 * time.Second)

	// 在定时器过期前重置它，将过期时间改为1秒
	go func() {
		time.Sleep(2 * time.Second)
		if !t.Reset(time.Second) {
			fmt.Println("Timer has expired or was stopped", time.Now())
		} else {
			fmt.Println("Timer was active and has been reset", time.Now())
		}
	}()

	// 从定时器的通道中读取值，等待定时器过期
	select {
	case <-t.C:
		fmt.Println("Timer expired", time.Now())
	case <-time.After(5 * time.Second):
		fmt.Println("Timer did not expire within 5 seconds", time.Now())
	}

	// 停止定时器
	t.Stop()
}