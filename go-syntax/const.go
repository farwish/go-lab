package main

import (
    "fmt"
    "unsafe"
)

// 显式类型定义
const A string = "AA"
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

func main() {
    fmt.Println(A)
    fmt.Println(B)
    fmt.Println(Male)
    fmt.Println(unsafe.Sizeof(A)) // "AA"=16(byte)

    fmt.Println(a) // 0
    fmt.Println(b) // 1
    fmt.Println(c) // 2

    fmt.Println(d) // 0
    fmt.Println(e) // 1
    fmt.Println(f) // 2
}
