package main

// #include <stdio.h>
// #include <stdlib.h>
/*
   void print_fun(char *s) {
       printf("print used by C language: %s\n", s);
   }
*/
import "C"

// Go 代码直接写入 C 代码。
// 注意：c 代码和 import "C" ​​语句之间不能有空行！
func main() {
	cStr := C.CString("hello world !")
    C.print_fun(cStr)
}