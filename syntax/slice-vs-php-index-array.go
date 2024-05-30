package main

import "fmt"

func main() {
	// 在Go语言中，数组（array）是固定长度的数据结构，一旦声明其大小就不能改变。
	// 因此，你不能直接向一个已声明的数组中添加元素。但是，如果你有一个可变大小的需求，你应该使用切片（slice）而不是数组。
	// 效果等同于php的数组 $slice = [];
	var slice []interface{}   // 或者指定切片值类型，如 []int、[]string
	fmt.Println(slice)
	slice = append(slice, 1, "a", "b", "c")
	fmt.Println(slice)

	// 或者直接初始化
	// var s2 := []interface{}{1, "a", "b", "c"}

	// 【遍历切片】

	// 遍历切片: for循环
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// 遍历切片：使用for...range关键字, 同时获取索引和值
	for i, value := range slice {
		fmt.Println(i, value)
	}
	// 遍历切片：使用range关键字, 仅获取值，不关心索引
	for _, value := range slice {
		fmt.Println(value)
	}

	// 在 Go 语言中，使用 range 遍历切片和数组的效率与直接通过索引访问相差不大。
	// Go 编译器会对 range 进行优化，因此性能上通常不会成为瓶颈。
	// 然而，从代码简洁性和可读性角度看，使用 range 更为推荐，因为它能减少手动处理索引的操作，
	// 并且对迭代任何序列类型的接口都保持一致。

	// 【删除切片中某个索引的值】
	indexToRemove := 2
	// 使用三个点操作符将一个切片的所有元素一次性追加到另一个切片后面
	slice = append(slice[:indexToRemove], slice[indexToRemove+1:]...)
	fmt.Println(slice) 	// [1 a c]

}