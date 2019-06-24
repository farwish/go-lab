package main

import (
    "fmt"
)

func main() {
    arr := [...]int {2, 4, 6, 8, 10}

    // 切片初始化
    s := []int {1, 2, 3}
    fmt.Printf("s slice type is %T\n", s) // []int
    printSlice(s) // len=3, cap=3, slice=[1 2 3]

    // make 初始化切片
    ss := make([]int, 3)
    fmt.Printf("ss slice type is %T\n", ss) // []int
    printSlice(ss) // len=3, cap=3, slice=[0 0 0]

    // 初始化一个长度和ss相等的，容量2倍的切片
    ssDouble := make([]int, len(ss), cap(ss)*2)
    printSlice(ssDouble)

    // 通过数组初始化切片sss
    sss := arr[0:4]
    fmt.Printf("ss slice type is %T\n", sss) // []int
    printSlice(sss) // len=3, cap=3, slice=[2 4 6 8]

    // 通过切片初始化切片ssss
    ssss := sss[:]
    fmt.Printf("ssss slice type is %T\n", ssss) // []int
    printSlice(ssss) // len=3, cap=3, slice=[2 4 6 8]

    // 空切片
    var numbers []int
    printSlice(numbers) // len=0, cap=0, slice=[]
    if (numbers == nil) {
        fmt.Println("numbers slice is nil")
    }

    // 切片截取: startIndex 包含, endIndex 不包含
    // sss = [2 4 6 8], sss[1:2] = [4]
    fmt.Printf("sss = %v, sss[1:2] = %v\n", sss, sss[1:2])
    // sss = [2 4 6 8], sss[1:] = [4 6 8]
    fmt.Printf("sss = %v, sss[1:] = %v\n", sss, sss[1:])
    // sss = [2 4 6 8], sss[:2] = [2 4]
    fmt.Printf("sss = %v, sss[:2] = %v\n", sss, sss[:2])

    // 增加切片容量，append() 和 copy()
    ss = append(ss, 0)
    ss = append(ss, 1)
    ss = append(ss, 2, 3, 4)
    // len=8, cap=12, slice=[0 0 0 0 1 2 3 4]
    printSlice(ss)

    // 拷贝 ss 的内容到 ssDouble
    copy(ssDouble, sss)
    // len=3, cap=6, slice=[2 4 6]
    printSlice(ssDouble)
}

// 长度和容量
func printSlice(x []int) {
    fmt.Printf("len=%d, cap=%d, slice=%v\n", len(x), cap(x), x)
}
