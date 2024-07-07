package main

import (
	"fmt"
	"net"
)

func main() {
	
	ipv4Private1 := net.ParseIP("10.255.0.0")
	ipv4Private2 := net.ParseIP("172.16.0.0")
	ipv4Private3 := net.ParseIP("192.168.0.0")
	
	ipv4Public := net.ParseIP("11.0.0.0")

	ipv6Private := net.ParseIP("fc00::")
	ipv6Public := net.ParseIP("fe00::")

	
	fmt.Println(ipv4Private1.IsPrivate())	// true
	fmt.Println(ipv4Private2.IsPrivate()) 	// true
	fmt.Println(ipv4Private3.IsPrivate())	// true

	fmt.Println(ipv4Public.IsPrivate()) 	// false

	fmt.Println(ipv6Private.IsPrivate()) 	// true
	fmt.Println(ipv6Public.IsPrivate())		// false
}