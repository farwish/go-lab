package main

import (
    "fmt"
    "time"
)

var quit chan int = make(chan int)

func main() {
    var j int = 0
    startTime := time.Now().Unix()

    for j = 0; j < 5; j++ {
        // 串行, 总耗时10s
        //say("world")

        // 并发调用入队, 总耗时2s
        go say("world")
    }

    // 出队
    for j = 0; j < 5; j++ {
        <- quit
    }

    endTime := time.Now().Unix()

    fmt.Printf("total cost %d seconds\n", endTime - startTime)
}

// 耗时2s
func say(s string) {
    time.Sleep(2 * time.Second)
    fmt.Println(s)

    // 入队
    quit <- 0
}
