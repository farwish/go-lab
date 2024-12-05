package main

import (
	"log"

	"github.com/robfig/cron/v3"
)

func main() {
    // 创建一个默认的cron对象
    c := cron.New()

    // 添加任务
    c.AddFunc("1/* * * * *", func() { log.Println("Every minute") })
    // c.AddFunc("30 * * * *", func() { fmt.Println("Every hour on the half hour") })
    // c.AddFunc("30 3-6,20-23 * * *", func() { fmt.Println(".. in the range 3-6am, 8-11pm") })
    // c.AddFunc("@hourly",      func() { fmt.Println("Every hour, starting an hour from now") })
    // c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty, starting an hour thirty from now") })

    //开始执行任务
    c.Start()

    //阻塞
    select {}
}