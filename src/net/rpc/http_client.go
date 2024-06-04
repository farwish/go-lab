package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	X, Y int
}
type Reply struct {
	Sum int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := Args{X: 7, Y: 8}
	var reply Reply
	err = client.Call("Arith.Multiply", &args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	log.Printf("Arith: %d * %d = %d", args.X, args.Y, reply.Sum)
}
