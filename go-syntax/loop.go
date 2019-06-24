package main

import (
    "fmt"
    "time"
    "unicode/utf8"
)

func main() {
    // 补充内容: 获取utf8字符长度
    greeting := "你好"
    fmt.Printf("'%s' 长度为 %d\n", greeting, utf8.RuneCountInString(greeting));

    // for.. range 遍历字符串，输出字符
    city := "乌鲁木齐"
    for _, value := range city {
        v := fmt.Sprintf("%c", value)
        fmt.Println(v)
    }

    i, j, k := 0, 0, 0

    // 和 C 的 while 一样
    for true {
        // time.Nanosecond  1纳秒
        // time.Microsecond 1微妙
        // time.Millisecond 1毫秒
        // time.Second      1秒

        fmt.Println("infinite loop1")
        time.Sleep(1 * time.Millisecond)

        i++;
        if i == 10 {
            break;
        }
    }

    // 和 C 的 for(;;) 一样
    for {
        fmt.Println("infinite loop2")
        time.Sleep(1 * time.Millisecond)

        j++;
        if j == 10 {
            break;
        }
    }

    // 下面使用 continue 和 goto 都可以
    for {
        START:

        k++
        if k == 10 {
            break;
        }

        if (k % 2 == 0) {
            goto START
        }
        fmt.Println(k)
    }
}
