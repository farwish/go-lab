package main

import "fmt"

func main() {
	// 在Go语言中，并没有明确区分“索引map”和“关联map”的概念。
	// Go语言中的映射（map）是一种键值对的数据结构，它的键可以是任何可比较的类型，通常包括整数、字符串等。
	// 但无论键是什么类型，它们都是通过键的哈希值进行查找和存储的，因此并不强调键是否连续或有序。

	assocMap := map[string]interface{}{
		"a": 1,
		"b": "2",
		"c": "c",
	}
	assocMap["d"] = 4;
	fmt.Println(assocMap["b"])	// 2
	// 删除键值
	delete(assocMap, "b")
	fmt.Println(assocMap)		// map[a:1 c:c d:4]


	anotherMap := map[int]interface{}{
		1: "A",
		2: "B",
		3: "C",
	}
	// 插入键值
    anotherMap[4] = "D"
	fmt.Println(anotherMap) 	// map[1:A 2:B 3:C 4:D]
}