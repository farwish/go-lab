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
    var name string             // 字符串型。string类型是包含8位字节的集合，通常表示UTF-8编码的文本。字符串可以为空，但不能为nil。字符串值是不可变的。
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
    //var M byte        // byte 是 uint8 的【别名】, See: $GOROOT/src/builtin/builtin.go
    //var N rune        // rune 是 int32 的【别名】, See: $GOROOT/src/builtin/builtin.go
    //var O uint        // 32 位无符号整数。它至少有32位，它是一个独立的类型，而不是其他类型（如uint32）的别名。
    //var P int         // 32 位有符号整数。它至少有32位，它是一个独立的类型，而不是int32等类型的别名。
    //var Q uintptr     // uintptr是一个足够大的整型，能够存储任何指针的位模式。

    // 任意类型
    // var R interface{}
    // var R2 any      // any 是 interface{} 的【别名】, See: $GOROOT/src/builtin/builtin.go
}
