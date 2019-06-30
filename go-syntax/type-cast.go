package main

import (
    "fmt"
    "strconv"
)

func main() {
    sum := 17
    count := 2
    var avg float32

    // 整型转为浮点型
    avg = float32(sum) / float32(count)
    fmt.Printf("avg = %f\n", avg) // 8.500000

    sumStr := strconv.Itoa(sum)
    fmt.Printf("%s\n", sumStr)
}
