package main

import (
    "fmt"
)

func main() {
    var a int = 10
    var ip *int = &a
    var pointer *int

    // 变量的存储地址
    fmt.Printf("address of a is %p\n", &a)

    // 指针变量的存储地址
    fmt.Printf("address of ip is %p\n", ip)

    // 用指针访问值
    fmt.Printf("ip value is %d\n", *ip)

    // 空指针值为 nil, 但不能(*ptr)对空指针解引用
    if (pointer == nil) {
        fmt.Println("ptr is nil")
    } else {
        fmt.Println("ptr is not nil")
    }

    // 数组的地址不等于数组第一个元素的地址
    arr := []int {1, 2, 3}
    fmt.Printf("arr    address is %p\n", &arr)
    fmt.Printf("arr[0] address is %p\n", &arr[0])
    fmt.Printf("arr[1] address is %p\n", &arr[1])
    fmt.Printf("arr[2] address is %p\n", &arr[2])

    // 指针数组
    var p [3]*int
    var b = []int {1, 2, 3}
    var i int
    for i = 0; i < 3; i++ {
        p[i] = &b[i] // 整数的地址存入指针数组
    }

    for i = 0; i < 3; i++ {
        fmt.Printf("p[%d] = %d\n", i, *p[i]) // 解引用
    }

    // 指针的指针
    var c int = 300
    var ptr *int
    var pptr **int

    ptr = &c
    pptr = &ptr // 指向指针的地址

    fmt.Printf("ptr = %d\n", *ptr)
    fmt.Printf("pptr = %d\n", **pptr) // 解引用两次

    // 指针作为函数参数
    x, y := 5, 6
    swap(&x, &y)
    fmt.Printf("x = %d\n", x) // 6
    fmt.Printf("y = %d\n", y) // 5

    y, x = x, y
    fmt.Printf("x = %d\n", x) // 5
    fmt.Printf("y = %d\n", y) // 6
}

func swap(x *int, y *int) {
    *x, *y = *y, *x

    // 中间变量
    //var temp int
    //temp = *x
    //*x = *y         // 赋值
    //*y = temp       // 赋值
}
