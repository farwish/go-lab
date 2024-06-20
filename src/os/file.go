package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var filepath = "D:\\x.go"
	
	// os.Open() 该函数用于打开指定名称的文件以进行读取。
	// 成功打开后，可以使用返回的文件对象进行读取操作，其关联的文件描述符以 O_RDONLY 模式打开。
	// 如果出现错误，将返回类型为 *PathError 的错误。实际上，该函数是通过调用 OpenFile 函数来实现的。
	if file, err := os.Open(filepath); err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()

		// x.Name()
		fmt.Println(file.Name())

		// x.Read() 该函数是一个File类型的成员函数，用于从文件中读取最多 len(b) 字节的数据，并将这些数据存储在 b 中。
		// 函数返回实际读取的字节数以及可能遇到的错误。
		// 当到达文件末尾时，函数返回 0 和 io.EOF 错误。在执行读取操作之前，函数会先检查文件是否有效，如果无效则返回错误。
		// 实际的读取操作由 f.read(b) 完成，函数会将读取的字节数和错误进行处理并返回。
		buffer := make([]byte, 5)
		if n, err := file.Read(buffer); err != nil && err != io.EOF {
			fmt.Println(err)
		} else {
			// 打印读取到的1024长度内容
			fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))
		}
	}


	// UserConfigDir() 函数返回用于存储用户特定配置数据的默认根目录。根据不同的操作系统，它返回不同的路径。
	// 在 Unix 系统上，它返回 $XDG_CONFIG_HOME 环境变量指定的路径，如果该变量为空，则返回 $HOME/.config。
	// 在 Darwin（macOS）系统上，它返回 $HOME/Library/Application Support。
	// 在 Windows 系统上，它返回 %AppData%。
	// 在 Plan 9 系统上，它返回 $home/lib。
	// 如果无法确定位置（例如，$HOME 未定义），则会返回错误。
	if ucd, err := os.UserConfigDir(); err == nil {
		fmt.Println(ucd)	// C:\Users\Administrator\AppData\Roaming
	}

	// UserHomeDir() 函数返回当前用户的主目录。根据不同的操作系统，它返回不同的路径。
	// 在 Unix（包括macOS）系统上，它返回 $HOME 环境变量。
	// 在 Windows 系统上，它返回 %USERPROFILE%。
	// 在 Plan 9系统上，它返回 $home 环境变量。
	// 如果相应的环境变量未设置，函数将返回一个默认值或错误。
	if udh, err := os.UserHomeDir(); err == nil {
		fmt.Println(udh)	// C:\Users\Administrator
	}

	// os.ReadFile() 用于读取指定文件名的文件内容，并返回读取到的字节数据和可能出现的错误。
	if b, err := os.ReadFile(filepath); err == nil {
		fmt.Println(string(b))
	}
}