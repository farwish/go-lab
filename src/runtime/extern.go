package main

import "runtime"

func main() {
	// 命令 "go tool dist list" 可获取支持的所有 操作系统(GOOS) 和 架构(GOARCH) 的组合列表。
	println(runtime.GOOS)		// windows
	println(runtime.GOARCH)		// amd64

	println(runtime.Version())	// go1.22.2
	println(runtime.GOROOT())	// D:\Go

	// 该函数用于获取调用方的文件名和行号信息。参数skip表示要跳过的栈帧数，0表示调用caller的函数。返回值分别表示对应的调用程序计数器、文件名和文件中的行号。布尔值ok为false表示无法获取该信息。
	// runtime.Caller()

	// 该函数用于获取调用者goroutine的堆栈上函数调用的返回程序计数器，并将其填充到切片pc中。
	// skip参数表示在记录到pc之前要跳过的堆栈帧数，其中0表示Callers自身的帧，1表示Callers的调用者帧。
	// 函数返回写入到pc中的条目数。建议使用CallersFrames来将这些PC转换为符号信息，如函数名称和行号，并且不建议直接迭代返回的PC切片或使用FuncForPC来获取任何返回的PC的信息，因为这些无法考虑内联或返回程序计数器的调整。
	// runtime.Callers()
}