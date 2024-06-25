package main

import "embed"

//go:embed hello.txt
var f embed.FS

// content holds our static web server content.
//
//go:embed image/* template/*
//go:embed html/index.html
var content embed.FS

func main() {
	// go:embed cannot apply to var inside func

	data, _ := f.ReadFile("hello.txt")
	println(string(data)) // hi there

	html, _ := content.ReadFile("html/index.html")
	one, _ := content.ReadFile("image/test1.txt")
	two, _ := content.ReadFile("template/test2.txt")
	println(string(html)) // this is index.html content.
	println(string(one))  // this is test1.txt content.
	println(string(two))  // this is test2.txt content.
}
