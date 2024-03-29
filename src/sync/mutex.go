package main

import (
	"fmt"
	"sync"
)

var count int
var mtx sync.Mutex

// By GPT：
// 在这个例子中，我们使用 sync.Mutex 保护了 count 变量，
// 确保在多个并发 Goroutine 同时对其进行递增操作时不会出现竞态条件。
// 每个 Goroutine 在修改 count 值之前都会先获取锁，修改后再释放锁，这样就能确保并发安全。

// 死锁：
// 下面的例子中不能在 mtx.Lock() 之后加 defer mtx.Unlock()，因为整个函数执行完才会执行 defer，而不是在一个循环中; 同样，也不能加在函数里面的最上方。
// 执行耗时任务时 其它goroutine无法获取到锁，count++后直接解锁够了。
func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		mtx.Lock()
		
		// 假设此处为复杂逻辑，可能 panic
        complexLogicThatMightPanic()
		
		count++
		fmt.Printf("Count: %d (from Goroutine)\n", count)

		mtx.Unlock()
	}
}

func complexLogicThatMightPanic() {
	// 这里是可能导致 panic 的复杂逻辑
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go increment(&wg)
	go increment(&wg)

	wg.Wait()
	fmt.Println("Final Count:", count)
}