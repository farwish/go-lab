package main

import (
    "fmt"
)

func main() {
    nums := []int {2, 3, 4} // 切片(没有长度和省略号)
    sum := 0

    // range 用于 for 循环中迭代数组/切片/通道/集合
    for _, num := range nums {
        // 如果for后只带一个变量，那么是键
        sum += num
    }
    fmt.Println("sum =", sum) // sum = 9

    // range 遍历 map
    kvMap := map[string]string {"a": "apple", "b": "banana"}
    for k, v := range kvMap {
        fmt.Printf("%s:%s\n", k, v)
    }

    // 遍历字符串
    for i, c := range "golang" {
        fmt.Printf("%d -> %c\n", i, c)
    }
}

