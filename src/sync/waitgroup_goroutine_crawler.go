package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func fetch(url string, ch chan<- string, wg *sync.WaitGroup) {
    defer wg.Done()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprintf("Error fetching %s: %v", url, err)
        return
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    ch <- fmt.Sprintf("Fetched %s: %s", url, body)
}

func main() {
    urls := []string{
        "http://example.com",
        "http://example.org",
        "http://example.net",
    }

    var wg sync.WaitGroup
    ch := make(chan string)

    for _, url := range urls {
        wg.Add(1)
        go fetch(url, ch, &wg)
    }

    go func() {
        wg.Wait()
        close(ch)
    }()

    for msg := range ch {
        fmt.Println(msg)
    }
}