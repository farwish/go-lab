package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('+')
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}
		fmt.Println("Received: ", message)
		conn.Write([]byte(message)) // Echo back the message
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			continue
		}
		go handleConnection(conn) // Handle each connection in a separate goroutine
	}
}