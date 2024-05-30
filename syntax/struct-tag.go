package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	Age int
	FirstName string
	LastName string `json:"last_name"`
}

type GoodMan struct {
	*Person
	sex int
}

func main() {
	p := Person{30, "Bob", "Bran"}

	log.Println(p)	// {30 Bob Bran}

	p.Age = 20
	log.Println(p)  // {20 Bob Bran}

	// 结构体加了第三个位置的Tag，在转换指定类型时，key 会使用指定的名字

	j, _ := json.Marshal(p)
	log.Println(string(j))	// {"Age":20,"FirstName":"Bob","last_name":"Bran"}
}

