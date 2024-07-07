package main

import (
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	g := errgroup.Group{}

	// 限制在该 errgroup 中,同时运行的goroutines数量最多为 n 个
	g.SetLimit(2)

	log.Println("/Start/")

	// 第一个goroutines中执行
	g.Go(DoSomething)
	// 尝试第二个goroutines
	r2 := g.TryGo(DoSomething)
	log.Println(r2)			// true
	// 尝试第三个goroutines
	r3 := g.TryGo(DoSomething)
	log.Println(r3)			// false，TryGo检测通道已满会直接返回false

	g.Go(DoSomething)
	g.Go(DoSomething)
	g.Go(DoSomething)
	g.Go(DoSomething)

	if e := g.Wait(); e != nil {
		log.Println("Some error:", e.Error())
	}
}

func DoSomething() error {
	time.Sleep(2 * time.Second)
	log.Println("done -",time.Now().Unix());
	return nil
}