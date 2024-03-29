# sync

【By GPT】

Go 语言标准库中的 sync 包提供了多种同步原语，用于控制多个goroutine间的协作和同步。
以下是sync包中最常用的一些类型及其基本用法：

Mutex（互斥锁）

```go
import "sync"

// 声明一个互斥锁
var mu sync.Mutex

// 在需要同步的代码块前获取锁
mu.Lock()
// 修改共享资源
// ...
mu.Unlock() // 释放锁

// 可以使用defer来保证锁一定会被释放
mu.Lock()
defer mu.Unlock()
// 修改共享资源
```

RWMutex（读写锁）

```go
import "sync"

var rwMutex sync.RWMutex

// 对共享资源进行读取操作
rwMutex.RLock()
// 读取共享资源
rwMutex.RUnlock()

// 对共享资源进行写入操作
rwMutex.Lock()
// 修改共享资源
rwMutex.Unlock()
```

Once

```go
import "sync"

var once sync.Once
var somethingInit bool

func setup() {
    // 初始化动作
    somethingInit = true
}

func doSomething() {
    once.Do(setup)
    // 确保setup只会被执行一次
    if somethingInit {
        // 执行依赖于初始化的动作
    }
}
```

WaitGroup

```go
import "sync"

var wg sync.WaitGroup

func worker(id int) {
    defer wg.Done()
    // 执行耗时任务
}

func main() {
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go worker(i)
    }

    // 等待所有worker完成
    wg.Wait()
    // 继续执行主线程后续逻辑
}
```

Cond

```go
import "sync"

var cond = sync.NewCond(new(sync.Mutex))

func worker(cond *sync.Cond) {
    cond.L.Lock()
    for !conditionMet() {
        cond.Wait()
    }
    // 当条件满足时继续执行
    cond.L.Unlock()
}

func notifier(cond *sync.Cond) {
    // 更新条件
    changeCondition()
    cond.Signal() // 唤醒一个等待的goroutine
    // 或
    cond.Broadcast() // 唤醒所有等待的goroutine
}
```

Pool

sync.Pool 是 Go 语言的标准库提供的一种轻量级对象池实现，主要用于缓存和重用临时对象，从而减少内存分配和垃圾回收（GC）的压力。其设计目标是帮助开发者在短期内复用已有的对象而不是频繁地创建和销毁它们，尤其是在那些生命周期较短且创建成本较高的场景下。

```go
import "sync"

var pool = sync.Pool{
    New: func() interface{} { return new(MyType) },
}

func getFromPool() *MyType {
    obj := pool.Get().(*MyType)
    // 使用obj...
    // ...
    return obj
}

func putIntoPool(obj *MyType) {
    obj.Reset() // 清理对象状态
    pool.Put(obj)
}
```

以上是对sync包中部分类型的简单使用示例。
通过这些同步原语，可以有效地控制并发环境下的数据竞争，确保资源安全访问，并管理goroutine间的同步和通信。