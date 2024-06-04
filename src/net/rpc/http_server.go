package main

import (
	"log"
	"net/http"
	"net/rpc"
	"time"
)

// Args 和 Reply 是参数和返回值的结构体，根据实际需求定义
type Args struct {
	X, Y int
}
type Reply struct {
	Sum int
}

// Arith 是我们要导出并注册为RPC服务的类型
type Arith struct{}

// Multiply 方法会被注册为RPC方法，客户端可以调用
func (t *Arith) Multiply(args *Args, reply *Reply) error {
	time.Sleep(5 * time.Second)
	
	reply.Sum = args.X * args.Y
	log.Printf("Received multiply request: %+v", args)
	log.Printf("Sending reply: %+v", reply)
	return nil
}

func main() {
	arith := new(Arith)
	// 注册RPC服务
	rpc.Register(arith)

	// 设置监听端口
	rpc.HandleHTTP()

	log.Println("Starting RPC server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
