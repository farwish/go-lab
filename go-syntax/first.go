// 定义包名
package main

// 需要使用的包
import "fmt"

// go run first.go 编译并执行，不会留下可执行文件
// 或者
// go build -o binary/first first.go && ./binary/first

// 程序开始执行的函数, main 函数是每一个可执行文件必须包含的
func main () {
    // 当标识符(常量/变量/类型/函数名/结构字段)以一个大写字母开头，那么这种形式的对象就可以被外部包的代码导入使用
    // 标识符以小写字母开头，则对包外不可见，但在整个包的内部是可见并且可用的.
    // GO的字符串连接 +
    fmt.Println("Google " + "lang")

    var apple string
    var orange string
    var fruit string

    apple = "Apple"
    orange = "Orange"
    fruit = apple + orange;

    fmt.Println("Fruit " + fruit);
}

/*
行分割符
注释
标识符:命名变量,类型等程序实体
字符串连接
关键字
*/

// 关键字
/*
break, default, func, interface, select,
case, defer, go, map, struct,
chan, else, goto, package, switch
const, fallthrough, if, range, type,
continue, for, import, return, var,
*/

// 预定义标识符
/*
append, bool, byte, cap, close, complex, complex64, complex128, uint16
copy, false, float32, float64, imag, int, int8, int16, uint32
int32, int64, iota, len, make, new, nil, panic, uint64
print, println, real, recover, string, true, uint, uint8, uintptr
*/
