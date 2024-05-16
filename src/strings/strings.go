package main

import (
	"fmt"
	"strings"
)

func main() {
	// 包strings提供了用于操作UTF-8编码字符串的简单函数。
	// 有关Go中UTF-8字符串的更多信息，请参阅 https://blog.golang.org/strings。

	// 函数名称： EqualFold
	// 判断两个字符串 s 和 t 在简单 Unicode 大小写转换下是否相等。
	// 该函数会将字符串解析为 UTF-8 编码，然后进行大小写转换比较。
	// 这种比较方式是一种更通用的大小写不敏感形式。
	str1 := "HELLO"
	str2 := "hello"
	result := strings.EqualFold(str1, str2)
	fmt.Println(result) // Output: true
}
