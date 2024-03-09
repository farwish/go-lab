// Go语言的包说明文档，描述了sync包的基本功能和使用注意事项。
// 	sync包提供了基本的同步原语，如互斥锁，以及Once和WaitGroup类型。
// 	这些类型主要用于低级库例程，而高级同步最好通过通道和通信来实现。包含sync包中定义的类型的值不应被复制。

// 以下是 利用Go的通道进行高级同步的具体的例子。

package main

import "fmt"

// 假设我们有一个函数模拟从网络获取数据的任务
func fetchData(url string) (string, error) {
    // 模拟耗时操作（例如：实际的HTTP请求）
    // 这里简化为返回一个字符串和错误信息
    return fmt.Sprintf("Data for %s", url), nil
}

func main() {
    urls := []string{"url1", "url2", "url3"}

    // 创建一个带缓冲的通道，容量等于需要并发执行的任务数
    results := make(chan string, len(urls))

    // 对于urls中的每个URL，启动一个goroutine去获取数据
    for _, url := range urls {
        go func(url string) {
            data, err := fetchData(url)
            if err != nil {
                // 在实际应用中，可能需要处理错误
                panic(err)
            }
            // 当数据获取完成后，通过通道发送结果
            results <- data
        }(url)
    }

    // 遍历通道，接收所有goroutine的结果
    // 这个for循环会阻塞，直到results通道被关闭或填满
    for i := 0; i < len(urls); i++ {
        fmt.Println(<-results) // 输出获取到的数据
    }
}