package main

import (
    "fmt"
)

func main() {
    // 声明变量的一般形式是使用 var 关键字，可以一次声明多个变量
    // 1.指定变量类型，如果没有初始化，则变量默认为零值
    var country string = "cn"
    var city, town string = "sh", "mh"

    // 2.根据值自行判定变量类型
    var card = "card"
    var card1, card2 = 123, "card2"

    // 3.简短形式(推荐)，省略 var 不带声明,
    // 注意 := 左侧如果不是新的变量就产生编译错误

    // 不带声明格式的只能在函数体中出现
    // 等同于 var sex = 0
    sex := 0
    sex1, sex2 := 1, 2

    fmt.Println(country, city, town)
    fmt.Println(card, card1, card2)
    fmt.Println(sex, sex1, sex2)

    // 交换两个变量的值，类型必须一样
    sex, sex1, sex2 = sex2, sex1, sex
    fmt.Println(sex, sex1, sex2)

    var flag bool               // 布尔型
    var age int                 // 数字型
    var name string             // 字符串型
    var a *int                  // 指针类型(Pointer)
    var b []int                 // 数组类型
    var c map[string] int       // Map 集合类型(key类型和value类型),不初始化(make)为 nil map, 不能存放键值对
    var d chan int              // Channel 类型
    var e func(string) int      // 函数类型
    var f error                 // 接口类型(interface)
                                // 结构化类型(struct)
                                // 切片类型

    // 未初始化的变量为零值
    fmt.Println(flag)   // false
    fmt.Println(age)    // 0
    fmt.Println(name)   // "" 空字符串

    fmt.Println(a)      // nil
    fmt.Println(b)      // []
    fmt.Println(c)      // map[]
    fmt.Println(d)      // nil
    fmt.Println(e)      // nil
    fmt.Println(f)      // nil

    // 基于架构de数字型，无符号+有符号
    //var A uint8
    //var B uint16
    //var C uint32
    //var D uint64
    //var E int8
    //var F int16
    //var G int32
    //var H int64
    //// 基于架构de浮点型
    //var I float32
    //var J float64
    //var K complex64
    //var L complex128
    //// 其他数字类型
    //var M byte
    //var N rune
    //var O uint
    //var P int
    //var Q uintptr
}
