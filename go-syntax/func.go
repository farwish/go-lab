package main

import (
    "fmt"
    "math"
)

type Circle struct {
    radius float64
}

func main() {
    // 值传递
    fmt.Println(max(1,2)) // 2

    // 多返回值
    a, b := both("A", "B")
    fmt.Println(a, b)

    // 引用传递
    refPrintln(&a, &b)

    // 匿名函数变量赋值
    getSquare := func(x float64) float64 {
        return math.Sqrt(x)
    }
    fmt.Println(getSquare(9)) // 3

    // 闭包
    fmt.Println(getSequence()()) // 1
    fmt.Println(getSequence()()) // 1
    // number为一个函数
    number := getSequence()
    fmt.Println(number())   // 1
    fmt.Println(number())   // 2
    fmt.Println(number())   // 3

    // 函数方法
    var c1 Circle
    c1.radius = 10.00
    fmt.Println("圆的面积 = ", c1.getArea())
}

// func 声明函数, 参数列表和返回类型可选, 函数体
func max(num1, num2 int) int {
    var result int
    if (num1 > num2) {
        result = num1
    } else {
        result = num2
    }
    return result
}

// 多返回值
func both(a, b string) (string, string) {
    return a, b
}

// 引用传递
func refPrintln(x *string, y *string) {
    fmt.Println(*x, *y);
}

// 返回函数类型
func getSequence() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

// getArea 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
    return 3.14 * c.radius * c.radius
}
