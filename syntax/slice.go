package main

import (
	"fmt"
	"slices"
	"unsafe"
)

// 与数组的区别：
//  1.数组有固定大小，不指定长度也会自动推断出，不可改变
//  切片没有固定大小，容量可以扩大
//  （ 从 %T 打印出的类型可以看出，数组是如 [5]int，切片是 []int ）
//  2.作为函数参数传递时，数组需要指名长度，切片不需要

func main() {
    arr := [...]int {2, 4, 6, 8, 10}

    // 切片初始化
    s := []int {1, 2, 3}
    fmt.Printf("s slice type is %T\n", s) // []int
    printSlice(s) // len=3, cap=3, slice=[1 2 3]

    fmt.Printf("arr     type is %T\n", arr)   // [5]int
    printArr(arr)

    // make 初始化切片, 必须指定长度，元素值会填充为零值，不会是空切片
    // 只是声明了就是空切片
    ss := make([]int, 3)
    fmt.Printf("ss slice type is %T\n", ss) // []int
    printSlice(ss) // len=3, cap=3, slice=[0 0 0]

    // 初始化一个长度和ss相等的，容量2倍的切片
    ssDouble := make([]int, len(ss), cap(ss)*2)
    printSlice(ssDouble) // len=3, cap=6, slice=[0 0 0]

    // 通过数组初始化切片sss
    sss := arr[0:4]
    fmt.Printf("sss slice type is %T\n", sss) // []int
    printSlice(sss) // len=3, cap=3, slice=[2 4 6 8]

    // 通过切片初始化切片ssss
    ssss := sss[:]
    fmt.Printf("ssss slice type is %T\n", ssss) // []int
    printSlice(ssss) // len=3, cap=3, slice=[2 4 6 8]

    // 空切片
    var numbers []int
    printSlice(numbers) // len=0, cap=0, slice=[]
    if (numbers == nil) {
        fmt.Println("numbers slice = nil")
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
    printSlice(ss)
    ss = append(ss, 2, 3, 4)
    // len=8, cap=12, slice=[0 0 0 0 1 2 3 4]
    printSlice(ss)

    // 拷贝 ss 的内容到 ssDouble
    copy(ssDouble, sss)
    fmt.Println("After copy:")
    // len=3, cap=6, slice=[2 4 6]
    printSlice(ssDouble)

    // Go1.22: slices.Concat 返回包含两个切片内容的新切片。
    printSlice(slices.Concat(ssDouble, sss))
}

// 长度和容量
func printSlice(x []int) {
    fmt.Printf("元素长度=%d, 容量=%d, 切片=%v, 地址=%v\n", len(x), cap(x), x, unsafe.Pointer(&x))
}

func printArr(x [5]int) {
    fmt.Printf("arr 地址=%v\n", unsafe.Pointer(&x));
}
