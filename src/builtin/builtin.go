package main

import (
	"fmt"
	"reflect"
)

func main() {
	// append(slice []Type, elems ...Type) []Type
	// copy(dst, src []Type) int
	test := []string{"test"}
	fmt.Println(test)		// [test]
	
	src := []string{"hello", "world"}
	dst := make([]string, len(src)) // 包含两个空元素的dst
	dst = append(dst, "AA")
	fmt.Println(dst)		// [  AA]
	n := copy(dst, src)
	fmt.Println(dst)		// [hello world AA]
	fmt.Println(n)			// 2

	num := copy([]byte("hello your "), "world")
	fmt.Println(num)		// 5

	fmt.Println("-------------------------------")

	// delete(m map[Type]Type1, key Type)
	// 		从 map 中删除指定key的元素
	// len(v Type) int 
	// 		适用于 Array, Pointer to array, Slice, String, Channel
	// cap(v Type) int
	// 		对于 Array, Pointer to array 含义和 len()相同, Slice 指可以达到的最大长度, Channel 指缓冲区的容量
	// make() 分配和初始化这些类型的对象, slice, map, or chan (only).
	// 			make初始化 slice 时, size指定长度，可提供第二个参数指定容量、且必须不小于长度
	// 			make初始化 map 可省略size，见 builtin.go 注释
	// 			make初始化 channel,size为0或者忽略时，则通道是无缓冲的。 
	m := make(map[string]int)  
	v := reflect.ValueOf(m)
	fmt.Println("Type   of m is:", reflect.TypeOf(m)) // map[string]int
	fmt.Println("Kind   of m is:", v.Kind()) 	// map, Kind() 可以获得变量的基本类型分类, map
	fmt.Println("Value  of m is:", v)			// map[]
	fmt.Println("length of m is:", len(m))		// 0

	m["a"] = 1
	m["b"] = 2
	delete(m, "a")

	v = reflect.ValueOf(m)
	fmt.Println("Value  of m is:", v)			// map[b:2]
	fmt.Println("length of m is:", len(m))		// 1

	fmt.Println("-------------------------------")

	// new(Type) *Type
	// 		创建一个指向类型为 Type 的变量的指针
    i := new(int)
    *i = 42
    fmt.Println(*i) 	// 42

    s := new(string)
    *s = "Hello, World!"
    fmt.Println(*s) 	// Hello, World!

	fmt.Println("-------------------------------")

	// clear[T ~[]Type | ~map[Type]Type1](t T)
	// 		清理 "map" 和 "slice", 
	// 		map是清理掉元素, 得到空的 map.
	// 		slice 是把所有元素清理为 zero 值.
	clear(test)
	clear(dst)
	clear(m)
	fmt.Println(test)	// []
	fmt.Println(dst)	// [  ],由于有元素占位,所以是几个零值的元素
	fmt.Println(m)		// map[]

	fmt.Println("-------------------------------")

	// close(c chan<- Type)
	// 		关闭一个通道，该通道可以是双向的或只发送的。
	// 		关闭操作应该仅由发送方执行，而不是接收方。
	// 		关闭通道的效果是在发送最后一个值后停止通道的使用。
	// 		在从关闭的通道c接收到最后一个值后，任何从此通道接收的操作都将立即成功，且返回通道元素的零值。
	// 		接收操作的形式为x, ok := <-c，当通道为空且已关闭时，ok将被设置为false。
	ch := make(chan int)
	go func () {
		// producer
		for i := 1; i < 5; i++ {
			ch <- i
		}
		close(ch) // 发送完数据后关闭通道
	}()

	// consumer once
	x, ok := <-ch
	fmt.Println(x)		// 1
	fmt.Println(ok)		// true

	// consumer
	for v := range ch { // 当通道关闭且无更多数据时，循环会自动结束
		fmt.Println(v)	// 2, 3, 4
	}

	x, ok = <-ch
	fmt.Println(x)		// 0,      在从关闭的通道c接收到最后一个值后，任何从此通道接收的操作都将立即成功，且返回通道元素的零值
	fmt.Println(ok)		// false,  当通道为空且已关闭时 ok将被设置为false。

	fmt.Println("-------------------------------")

	// panic(v any)
	// 		在当前的 goroutine 中停止正常的执行。
	// 		当一个函数F调用 panic 时，F的正常执行会立即停止。
	// 		F调用的任何被延迟的函数都会以通常的方式运行，然后F返回到它的调用者。
	// 		对于调用者G来说，对F的调用 then 行为就像调用 panic，终止G的执行并运行任何被延迟的函数。
	// 		这种终止顺序被称为 panic，并且可以通过内置函数 recover 进行控制。
	// 			从Go 1.21开始，如果 panic 的参数是一个nil接口值或一个未类型的 nil，会导致运行时错误（不同的panic）。
	// 		可以通过设置环境变量 GODEBUG=panicnil=1 来禁用运行时错误。

	// 注: 尽管panic提供了在Go中一种类似于其他语言异常处理的机制，
	// 		但根据Go的设计哲学，它并不鼓励滥用panic作为常规错误处理手段，而是推荐通过返回错误值进行控制流管理。
	// 		在实际生产环境中，应当尽量避免因预期之外的条件引发 panic，除非这些条件确实表示了无法继续执行程序的严重问题。

	// recover() any
	// 		它允许程序管理一个发生 panic 的 goroutine 的行为。
	// 		在一个延迟执行的函数（而不是由它调用的任何函数）内部执行 recover 调用会停止panic序列，
	// 		通过恢复正常的执行并检索传递给panic调用的错误值。如果在延迟函数外部调用recover，则不会停止.
}
