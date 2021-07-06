package main

import (
	"fmt"
	"time"
)

func main() {
	go producer("First ")
	go producer("Second")

	customer()
}

var channel = make(chan string)

func producer(flag string) {
	for {
		channel <- fmt.Sprintf("%s: %v", flag, time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(2 * time.Second)
	}
}

func customer() {
	for {
		message := <- channel
		fmt.Println(message)
	}
}
