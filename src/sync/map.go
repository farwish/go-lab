package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	/*
	 	定义了一个并发安全的Map数据结构，可以在多个 goroutine 之间共享和访问而无需额外的锁定或协调。
		它适用于缓存或在多个 goroutine 之间共享数据的场景。
		并发安全的 Map使用了原子操作和互斥锁来保证并发访问的安全性，并通过读写分离的策略来提高并发性能。

		[问题]
		m := new(sync.Map) 初始化后，能不使用m.Store 而是 m["key"] = 1 写入吗

		[答]
		不可以直接使用m["key"] = 1这种方式来写入sync.Map。
		因为sync.Map并没有直接实现索引赋值操作，它不支持像普通map那样的直接通过索引操作来设置或获取键值对。
		为了保证并发安全，sync.Map提供了特定的方法来进行读写操作，如Store, Load, LoadOrStore, Delete, 和 Range等。
		如果你想类似地使用索引操作，你需要通过 sync.Map 提供的方法来间接实现。
	*/
    m := new(sync.Map)
	var g errgroup.Group
    
	for i := 0; i < 10; i++ {
		g.Go(func () error {
			log.Printf("goroutine %d", i)
			m.Store(i, i)
			return nil
		})
	}

	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		g.Go(func () error {
			if val, ok := m.Load(i); ok {
				fmt.Println(val) // 安全的并发读取
			}
			return nil
		})
	}

	if err := g.Wait(); err == nil {
        fmt.Println("Successfully done.")
    }

}