package main

import (
	"fmt"
	"os"
)

// env.go 是Go语言标准库中的os包的一部分，提供了一些与操作系统环境变量相关的函数。
// 【仅影响调用该函数的进程本身，而不会影响操作系统级别的环境变量或其它正在运行的进程。】

// Expand 	 函数用于将字符串中的${var}或$var形式的变量扩展成对应的值。可以通过传入一个自定义的函数来指定变量的值如何获取。
// ExpandEnv 函数用于将字符串中的环境变量扩展成对应的值，使用当前环境变量的值来替换${var}或$var。
// Getenv 	 函数用于获取指定环境变量的值。
// LookupEnv 函数用于获取指定环境变量的值，并返回一个布尔值表示该变量是否被设置。
// Setenv 	 函数用于设置指定环境变量的值。
// Unsetenv  函数用于删除指定环境变量。
// Clearenv  函数用于清除所有环境变量。
// Environ 	 函数用于返回当前环境变量的列表，每个环境变量以"key=value"的形式表示。
// 这些函数提供了对操作系统环境变量的读取、设置和删除等操作，方便开发者在程序中处理环境变量相关的逻辑。

func main() {
	// environ := os.Environ()
	// for _, env := range environ {
	// 	fmt.Printf("%s\n", env)
	// }
	// GO111MODULE=on
	// GOPATH=D:\Go\bin
	// GOPROXY=https://proxy.golang.com.cn,direct
	// HOME=C:\Users\Administrator
	// HOMEDRIVE=C:
	// ...

	env := os.Getenv("HOME")
	fmt.Println(env)	// C:\Users\Administrator

	r, b := os.LookupEnv("HOME")
	fmt.Println(r, b)	// C:\Users\Administrator  true
	r2, b2 := os.LookupEnv("HOME2")
	fmt.Println(r2, b2)	//  false

	os.Setenv("HOME3", "C:\\Users\\Administrator")
	r3, b3 := os.LookupEnv("HOME3")
	fmt.Println(r3, b3)	// C:\Users\Administrator  true

	os.Unsetenv("HOME3")
	r4, b4 := os.LookupEnv("HOME3")
	fmt.Println(r4, b4) // false
}
