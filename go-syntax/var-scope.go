package main

import (
    "fmt"
)

// 声明全局变量
var a int

func main() {
    // 声明局部变量
    var a int = 1
    var b int = 10

    fmt.Println(a)          // 0
    fmt.Println(sum(a, b)) // 11
}

// 形式参数
func sum(a, b int) int {
    return a + b
}
