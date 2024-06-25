package main

import _ "embed"

//go:embed hello.txt
var s string

//go:embed hello.txt
var b []byte

func main() {
	// Note, go:embed cannot apply to var inside func

	println(s)         // hi there!
	println(string(b)) // hi there!
}
