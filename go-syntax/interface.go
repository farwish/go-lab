package main

import (
    "fmt"
)

// 定义接口 Phone, 方法 tel
type Phone interface {
    tel()
}

// 结构 -> 类
type NokiaPhone struct {
    name string
}

// 结构 -> 类
type IPhone struct {
    name string
}

// 参数非指针的函数方法
// (使用时需要先 new 这个结构)
func (nokia NokiaPhone) call() {
    fmt.Println("I am " + nokia.name)
}

func (iphone IPhone) call() {
    fmt.Println("I am " + iphone.name)
}

// 参数为指针的函数方法
func (nokia *NokiaPhone) callp() {
    fmt.Println("I am "+ nokia.name +" plus")
}

func (iphone *IPhone) callp() {
    fmt.Println("I am "+ iphone.name +" plus")
}

func main() {
    // 使用 new 实例化（因为函数方法的参数不是指针）
    nokia  := new(NokiaPhone)
    iphone := new(IPhone)
    nokia.name = "nokia-1"; iphone.name = "iphone-1"
    nokia.call()
    iphone.call()

    // 参数为指针的函数方法使用
    nokiaP  := NokiaPhone{name:"nokia"}
    iphoneP := IPhone{name:"iphone"}
    nokiaP.callp()
    iphoneP.callp()
}
