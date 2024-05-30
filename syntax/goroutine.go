package main

import (
    "fmt"
    "time"
)

var quit chan int = make(chan int)

func main() {
    fmt.Printf("quit chan type = %T\n", quit)
    fmt.Printf("quit chan value = %v\n", quit)

    var j int = 0
    startTime := time.Now().Unix()

    for j = 0; j < 5; j++ {
        // 串行, 总耗时10s
        //say(j, "world")

        // 并发调用入队, 总耗时2s
        go say(j, "world")
    }

    // 出队
    for j = 0; j < 5; j++ {
        <- quit
    }

    endTime := time.Now().Unix()

    fmt.Printf("Total cost %d seconds\n", endTime - startTime)
}

// 耗时2s
func say(idx int, str string) {
    time.Sleep(2 * time.Second)
    fmt.Printf("idx:%d -> %s\n", idx, str)

    // 入队
    quit <- 0
}
