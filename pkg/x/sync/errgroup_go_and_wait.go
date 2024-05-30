package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
    type headers map[string]string

    var g errgroup.Group            // 声明一个group实例
    var urls = []string{
        "http://www.offso.com/",
        "http://www.bing.com",
    }
    for _, url := range urls {      // 分别获取网站内容
        url := url                  // url是局部变量，for循环中对多个协程传递值时，需要重新进行赋值
        g.Go(func() error {         // group 的go方法，启一个协程去执行代码
            resp, err := MyGet(url, headers{    // Fetch the URL. Replace http.Get(url)
                "User-Agent": "Firefox",
            })
            if err == nil {
                resp.Body.Close()
            }
            return err
        })
    }
    if err := g.Wait(); err == nil {  // group 的wait方法，等待上面的 g.go的协程执行完成，并且可以接受错误
        fmt.Println("Successfully fetched all URLs.")
    }
}


func MyGet(url string, headers map[string]string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// 添加 Headers 到请求中
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return http.DefaultClient.Do(req)
}