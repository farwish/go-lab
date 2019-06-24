package main

import (
    "fmt"
)

func main() {
    // 声明一维数组
    var arr [5]string

    // 初始化数组: 含元素个数
    var balance = [5]int {1, 3, 5, 7, 9}
    balanceLen := len(balance)

    // 初始化数组: 不含元素个数,自动计算
    var balance2 = [...]int {2, 4, 6, 8, 10}
    balanceLen2 := len(balance2)

    fmt.Println(balanceLen)  // 5
    fmt.Println(balanceLen2) // 5

    // 设置元素值
    for i := 0; i < balanceLen; i++ {
        arr[i] = fmt.Sprintf("A%d", i)
    }

    // 获取元素值
    for j := 0; j < balanceLen; j++ {
        fmt.Printf("arr[%d]=%s\n", j, arr[j])
        // Output:
        //arr[0]=A0
        //arr[1]=A1
        //arr[2]=A2
        //arr[3]=A3
        //arr[4]=A4
    }

    // 二维数组
    a := [3][4]int {
        {0, 1, 2, 3},
        {4, 5, 6, 7},
        {8, 9, 10, 11},
    }
    alen := len(a)

    for k := 0; k < alen; k++ {
        for l := 0; l < len(a[k]); l++ {
            fmt.Printf("a[%d][%d]=%d\n", k, l, a[k][l])
            // Output:
            //a[0][0]=0
            //a[0][1]=1
            //a[0][2]=2
            //a[0][3]=3
            //a[1][0]=4
            //a[1][1]=5
            //a[1][2]=6
            //a[1][3]=7
            //a[2][0]=8
            //a[2][1]=9
            //a[2][2]=10
            //a[2][3]=11
        }
    }

    // 数组作为函数参数

    // 方式一
    //void myFunc (a [10]int) {

    //}

    // 方式二
    //void myFunc2 (b []int) {

    //}
}
