package main

import (
    "fmt"
)

// 接口, 方法call
type Phone interface {
    call()
}

// 结构 -> 类
type NokiaPhone struct {

}

// 结构 -> 类
type IPhone struct {

}

func (nokia NokiaPhone) call() {
    fmt.Println("I am nokia")
}

func (iphone IPhone) call() {
    fmt.Println("I am iphone")
}

func main() {
    nokia := new(NokiaPhone)
    iphone := new(IPhone)

    nokia.call()
    iphone.call()
}
