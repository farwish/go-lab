package main

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: ./sayhello/hello.a
#include <stdlib.h>
#include <sayhello/hello.h>
*/
import "C"
import "unsafe"

// Go 代码直接调用 C 文件。
// 注意：c 代码和 import "C" ​​语句之间不能有空行！
/*
# 编译 c 代码
$ cd sayhello/
$ gcc -c *.c           	# 生成.o文件
$ ar rs hello.a *.o 	# 生成.a文件，生成 .a 文件需要用到 ar 命令建立或修改备存文件
$ rm ./hello.o			# 删除已不需要的.o文件，linux下为 rm 命令，windows下为 del 命令

# 执行 go 代码
$ go run cog_file.go
*/
func main() {
	cStr := C.CString("hello world !") // golang 操作c 标准库中的CString函数
    C.print_fun1(cStr)                 // 调用C函数：print_fun 打印输出

    defer C.free(unsafe.Pointer(cStr)) // 因为 CSstring 这个函数没有对变量申请的空间进行内存释放
}