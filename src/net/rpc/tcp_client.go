package main

import (
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	X, Y int
}
type Reply struct {
	Sum int
}

func main() {
	// 直接通过TCP连接
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("net.DialTCP error:", err)
	}

	// 创建一个基于TCP连接的RPC客户端
	client := rpc.NewClient(conn)
	args := Args{X: 7, Y: 8}
	var reply Reply
	err = client.Call("Arith.Multiply", &args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	log.Printf("Arith: %d * %d = %d", args.X, args.Y, reply.Sum)

	// 关闭连接
	client.Close()
}
