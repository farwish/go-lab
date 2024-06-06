package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	serverAddress := "localhost:8080" // 服务器地址和端口
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to server...")

	// 发送消息
	message := "Hello+, Server!+"
	conn.Write([]byte(message))

	// 接收服务器响应
	reader := bufio.NewReader(conn)
	readLoop:
	for {
		response, err := reader.ReadString('+')
		if err != nil {
			if err == io.EOF {
				break readLoop  // 如果到达文件末尾，结束循环
			}
			fmt.Println("Error reading:", err.Error())
			return
		}

		fmt.Println("Server responded with:", response)
	}
}