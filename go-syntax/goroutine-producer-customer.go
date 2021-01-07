package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go producer("First")
	go producer("Second")

	customer()
}

var channel = make(chan string)

func producer(flag string) {
	for {
		channel <- fmt.Sprintf("%s: %v", flag, rand.Int31())
		time.Sleep(time.Second)
	}
}

func customer() {
	for {
		message := <- channel
		fmt.Println(message)
	}
}
