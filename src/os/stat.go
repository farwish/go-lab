package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Stat 函数用于返回指定文件的 FileInfo 对象，
	// 如果出错，则会返回一个 *PathError 类型的错误。
	// 在函数内部，它会调用 testlog.Stat 记录日志，然后再调用 statNolog 方法来实际获取文件信息。
	info, err := os.Stat("D:\\www\\go-lab\\src\\os\\env.go")
	if err != nil {
		fmt.Printf("【Error】%v\n", err)
		//【error】CreateFile /path/to/some/file: The system cannot find the path specified.
	} else {
		fmt.Printf("【File info】%+v\n", info) 
		//【File info】&{name:file FileAttributes:32 CreationTime:{LowDateTime:1050574194 HighDateTime:31105892} LastAccessTime:{LowDateTime:141280665 HighDateTime:31113719} LastWriteTime:{LowDateTime:1239981993 HighDateTime:31113717} FileSizeHigh:0 FileSizeLow:1898 ReparseTag:0 filetype:0 Mutex:{state:0 sema:0} path:/path/to/some/file vol:0 idxhi:0 idxlo:0 appendNameToPath:false}
	}

	fmt.Println();

	// 手动格式化输出 FileInfo 的关键属性
	printFileInfo(info)

	fmt.Println();

	// Lstat函数，也是用于返回指定文件的 FileInfo 对象，但它的行为略有不同。
	// 如果指定的文件是一个符号链接，Lstat 会返回这个符号链接的信息，而不会尝试去跟随这个链接。
	// 同样地，如果出错，则会返回一个*PathError类型的错误。
	// 在Windows系统中，如果文件是一个重解析点（如符号链接或挂载的文件夹），Lstat 会返回重解析点的信息，而不会尝试解析它。
	// 与 Stat 一样，Lstat 也会调用 testlog.Stat 记录日志，然后再调用 lstatNolog 方法来实际获取文件信息。
	linkInfo, err := os.Lstat("/path/to/some/file")
	if err != nil {
		fmt.Printf("【Error】%v\n", err)
	} else {
		fmt.Printf("【Symbolic link info】%+v\n", linkInfo)
	}
}

func printFileInfo(info os.FileInfo) {
	fmt.Printf("Name: %s\n", info.Name())	// Name: file
	fmt.Printf("Size: %d bytes\n", info.Size()) // Size: 1898 bytes
	fmt.Printf("Mode: %s\n", info.Mode().String()) // Mode: -rw-rw-rw-
	fmt.Printf("ModTime: %s\n", info.ModTime().Format(time.RFC3339)) // ModTime: 2024-06-01T09:01:33+08:00

	fmt.Printf("IsDir: %t\n", info.IsDir()) // IsDir: false
	fmt.Printf("Sys: %#v\n", info.Sys()) 
	// Sys可能包含平台特定信息。例如：
	// &syscall.Win32FileAttributeData{FileAttributes:0x20, CreationTime:syscall.Filetime{LowDateTime:0x3e9e7d72, HighDateTime:0x1daa364}, LastAccessTime:syscall.Filetime{LowDateTime:0xe41f90fd, HighDateTime:0x1dac1f8}, LastWriteTime:syscall.Filetime{LowDateTime:0x49e89fa9, HighDateTime:0x1dac1f5}, FileSizeHigh:0x0, FileSizeLow:0x76a}
}