package main

import (
	"fmt"
	"sync"
)

func process(data int, ch chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    result := data * 2
    ch <- result
}

func main() {
    data := []int{1, 2, 3, 4, 5}
    var wg sync.WaitGroup
    ch := make(chan int)

    for _, d := range data {
        wg.Add(1)
        go process(d, ch, &wg)
    }

    go func() {
        wg.Wait()
        close(ch)
    }()

    results := []int{}
    for result := range ch {
        results = append(results, result)
    }

    fmt.Println("Results:", results)
}