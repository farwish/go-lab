Go 语言的并发编程主要基于 Goroutines 和 Channels 这两个核心概念。以下是 Go 语言并发编程的主要范式：

### 1. Goroutines
Goroutines 是轻量级的线程，由 Go 运行时管理和调度。创建一个 Goroutine 非常简单，只需要在函数调用前加上 `go` 关键字即可。

#### 示例
```go
package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    say("hello")
}
```

### 2. Channels
Channels 用于 Goroutines 之间的通信和同步。通过 Channels 可以安全地传递数据，避免竞态条件。

#### 创建和使用 Channel
```go
package main

import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // 将结果发送到通道
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c // 从通道接收结果

    fmt.Println(x, y, x+y)
}
```

### 3. Select 语句
`select` 语句用于在多个通信操作中进行选择，当有多个通道操作时，`select` 会随机选择一个可以执行的操作。

#### 示例
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(time.Second * 1)
        c1 <- "one"
    }()

    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
```

### 4. 同步原语
Go 语言提供了多种同步原语，如 `sync.Mutex`、`sync.RWMutex`、`sync.WaitGroup` 等，用于更细粒度的控制并发。

#### 示例：使用 `sync.WaitGroup`
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // 任务完成时调用 Done
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1) // 增加 WaitGroup 计数
        go worker(i, &wg)
    }

    wg.Wait() // 等待所有任务完成
}
```

### 5. Context
`context` 包用于在 Goroutines 之间传递请求范围的值、取消信号和截止时间。这对于构建可取消和超时的并发程序非常有用。

#### 示例
```go
package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Worker received cancel signal")
            return
        default:
            fmt.Println("Worker is running")
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    go worker(ctx)

    time.Sleep(2 * time.Second)
    cancel() // 发送取消信号

    time.Sleep(1 * time.Second)
}
```

### 总结
Go 语言的并发编程范式主要围绕 Goroutines 和 Channels 展开，通过这些机制可以轻松实现高效、安全的并发程序。此外，`select` 语句、同步原语和 `context` 包提供了更多的控制和灵活性。