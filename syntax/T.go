package main

import (
	"fmt"
)

/*
Go语言中的泛型（Generics）是在Go 1.18版本中引入的一个重要特性，它允许你编写可重用的代码，而不需要为每种数据类型重复编写相同的逻辑。

func functionName[TypeParameters constraints](arguments) returnType {
    // 函数体
}

functionName: 	是泛型函数的名字。
TypeParameters: 是类型参数列表，放在函数名后的方括号中，用逗号分隔多个类型参数。每个类型参数代表一个待定的类型。
constraints: 	是类型约束，可选，用来限制类型参数可以被哪些具体类型实例化。它可以是一个预定义的约束如any（表示可以是任何类型），或者是自定义的接口约束。
arguments: 		是函数的参数列表，与非泛型函数相同。
returnType: 	是函数的返回类型，同样可以是泛型类型。


【自定义类型约束的示例】

type MyConstraint interface {
    // 这里列出约束条件，比如必须实现的方法
    SomeMethod()
}

func MyFunction[T MyConstraint](t T) T {
    // 函数体
}
*/

type comparable interface {
    ~int | ~float64 | ~string // 这里示例说明 T 必须是 int 或 float64 或 string 类型
}

// 类型约束，用来限制类型参数可以被哪些具体类型实例化。
// 它可以是一个预定义的约束如 any（表示可以是任何类型），或者是自定义的接口约束。
func PrintValue[T any](value T) {
	fmt.Println("Value is:", value)
}

// 类型约束，要求 T 实现了 comparable 接口，意味着 T 可以用于比较操作。
// 注意这里的类型约束不能用 any，因为 T 默认情况下并不保证是可以比较的,会编译报错。
// 比如尝试对不支持的操作（如比较、算术运算）使用该类型参数时，会在编译时遇到错误。
func Max[T comparable](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	PrintValue(42)	// 输出: Value is: 42

	fmt.Println(Max[int](1, 2))    // 输出: 2
	fmt.Println(Max[float64](3.14, 2.71)) // 输出: 3.14
	fmt.Println(Max[string]("apple", "banana")) // 输出: banana
}