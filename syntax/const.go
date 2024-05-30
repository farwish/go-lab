package main

import (
    "fmt"
    "unsafe"
)

// 显式类型定义
const A string = "AAAAA"
// 隐式类型定义
const B = "BB"
// 常量用作枚举
const (
    Unknown = 0
    Female = 1
    Male = 2
)
// iota 特殊常量，可以被编译器修改的常量
// (可以理解为 const 语句块的行索引)
// iota是Go的预定义标识符，用于在常量声明中提供递增的计数器。
//     每当遇到一个新的const块，iota的值都会重置为0。
const (
    a = iota
    b = iota
    c = iota
)

const (
    d = iota
    e
    f
)

const (
    g = 1 + iota
    h 
    i 
)

func main() {
    fmt.Println(A)
    fmt.Println(B)
    fmt.Println(Male)
    fmt.Println(unsafe.Sizeof(A)) // 16

    fmt.Println(a) // 0
    fmt.Println(b) // 1
    fmt.Println(c) // 2

    fmt.Println(d) // 0
    fmt.Println(e) // 1
    fmt.Println(f) // 2

    fmt.Println(g) // 1
    fmt.Println(h) // 2
    fmt.Println(i) // 3
}
