package main

import (
    "fmt"
)

//if,
//if.. else..,
//if.. else if..,
//switch, 默认每个case不会穿透执行, 强制执行后面一条 case 语句显示加 fallthrough
//type-switch, 判断某个interface变量中实际存储的类型
//select, 类似switch, 每个case必须是一个通信操作(发送或接收), select 随机执行一个可运行的case, 没有时运行default子句, 否则将阻塞, Go 不会对channel进行求值

func main() {
    // if.. else..
    var a int = 100;
    if a < 20 {
        fmt.Println("a < 20")
    } else if a <= 100 {
        fmt.Println("a <= 100")
    } else {
        fmt.Println("a > 20")
    }

    var grade, mark string
    score := 70

    // switch 每个 case 不会穿透执行
    switch score {
        case 50,60,70:
            grade = "C"
        case 80:
            grade = "B"
        case 90:
            grade = "A"
        default:
            grade = "D"
    }
    fmt.Println(grade) // C

    // switch 判断可以写在 case 里
    switch {
        case grade == "D", grade == "C":
            mark = "差"
        case grade == "B":
            mark = "中"
        case grade == "A":
            mark = "良"
    }
    fmt.Println(mark) // 差

    // 不带 fallthrough 输出 case 2 true
    // 带 fallthrough 强制执行后面一条 case 语句
    switch {
        case false:
            fmt.Println("case 1 false")
        case true:
            fmt.Println("case 2 true")
            fallthrough
        case false:
            fmt.Println("case 3 false")
        case false:
            fmt.Println("case 4 false")
    }

    var x interface {}

    switch x.(type) {
        case string:
            fmt.Println("x type is string")
        case int:
            fmt.Println("x type is int")
        case nil:
            // x type is <nil>
            fmt.Printf("x type is %T\n", x)
        default:
            fmt.Println("x type is unknown")
    }

}
