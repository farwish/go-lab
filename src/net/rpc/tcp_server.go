package main

import (
	"log"
	"net"
	"net/rpc"
	"time"
)

type Args struct {
	X, Y int
}
type Reply struct {
	Sum int
}

type Arith struct{}

func (a *Arith) Multiply(args *Args, reply *Reply) error {
	time.Sleep(5 * time.Second)

	reply.Sum = args.X * args.Y
	log.Printf("Received multiply request: %+v", args)
	log.Printf("Sending reply: %+v", reply)
	return nil
}

func main() {
	rpc.Register(new(Arith))

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error listening:", err.Error())
	}
	defer listener.Close()

	// see: $GOROOT/src/net/rpc/server_test.go
	rpc.Accept(listener)
}