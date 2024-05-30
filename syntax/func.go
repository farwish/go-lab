package main

import (
    "fmt"
    "math"
)

type Circle struct {
    radius float64
}

func main() {
    // 延迟执行函数
    defer func() {
        fmt.Println("main func execute finished")
    }()

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

    // 函数方法, 只能给 type 定义的类型用
    var c1 Circle
    c1.radius = 10.00
    fmt.Println("圆的面积 = ", c1.getArea())

    // 可变参数,只能是最后一个参数
    changeAble("ha", 1, 2, 3, 4)
}

// func 声明函数, 参数列表和返回类型可选, 函数体
func max(num1, num2 int) int {
    defer func() {
        fmt.Println("max func execute finished")
    }()

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
    defer func() {
        fmt.Println("both func execute finished")
    }()

    return a, b
}

// 引用传递(指针参数)
func refPrintln(x *string, y *string) {
    defer func() {
        fmt.Println("refPrintln func execute finished")
    }()

    fmt.Println(*x, *y);
}

// 返回函数类型
func getSequence() func() int {
    defer func() {
        fmt.Println("refPrintln func execute finished")
    }()

    i := 0
    return func() int {
        i++
        return i
    }
}

// getArea 属于 Circle 类型对象中的方法
// 只能给 type 定义的类型用
// 更多关于结构类型内容见 struct.go
func (c Circle) getArea() float64 {
    defer func() {
        fmt.Println("getArea func execute finished")
    }()

    return 3.14 * c.radius * c.radius
}

// 可变参数,只能是最后一个参数
func changeAble(str string, numbers... int) int {
    defer func() {
        fmt.Println("changeAble func execute finished")
    }()

    var sum int
    for _, value := range numbers {
        sum += value
    }
    fmt.Printf("str is %s, 可变参数列表 sum = %d\n", str, sum)
    return sum
}
