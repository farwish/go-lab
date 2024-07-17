package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"unsafe"
)

func main() {
	var negzero = math.Copysign(0, -1)
	var nonnilptr uintptr = uintptr(unsafe.Pointer(&negzero))
	var nilptr uintptr = uintptr(unsafe.Pointer(nil))
	var tests = []struct {
		x, y    any
		compare int
	}{
		{1, 2, -1},
		{1, 1, 0},
		{2, 1, +1},
		{"a", "aa", -1},
		{"a", "a", 0},
		{"aa", "a", +1},
		{1.0, 1.1, -1},
		{1.1, 1.1, 0},
		{1.1, 1.0, +1},
		{math.Inf(1), math.Inf(1), 0},
		{math.Inf(-1), math.Inf(-1), 0},
		{math.Inf(-1), 1.0, -1},
		{1.0, math.Inf(-1), +1},
		{math.Inf(1), 1.0, +1},
		{1.0, math.Inf(1), -1},
		{math.NaN(), math.NaN(), 0},
		{0.0, math.NaN(), +1},
		{math.NaN(), 0.0, -1},
		{math.NaN(), math.Inf(-1), -1},
		{math.Inf(-1), math.NaN(), +1},
		{0.0, 0.0, 0},
		{negzero, negzero, 0},
		{negzero, 0.0, 0},
		{0.0, negzero, 0},
		{negzero, 1.0, -1},
		{negzero, -1.0, +1},
		{nilptr, nonnilptr, -1},
		{nonnilptr, nilptr, 1},
		{nonnilptr, nonnilptr, 0},
	}

	// Less 函数用于比较两个类型为 T 的值 x 和 y 是否满足 x<y。
	// 其中，T 需要实现 Ordered 接口。对于浮点数类型，NaN 被认为小于任何非 NaN 值，-0.0 不小于（等于）0.0。
	// 函数通过判断 x 是否为 NaN 且 y 不是 NaN，或者 x 小于 y 来返回比较结果。

	for _, test := range tests {
		var b bool
		switch test.x.(type) {
		case int:
			b = cmp.Less(test.x.(int), test.y.(int))
		case string:
			b = cmp.Less(test.x.(string), test.y.(string))
		case float64:
			b = cmp.Less(test.x.(float64), test.y.(float64))
		case uintptr:
			b = cmp.Less(test.x.(uintptr), test.y.(uintptr))
		}
		if b != (test.compare < 0) {
			fmt.Printf("Less(%v, %v) == %t, want %t\n", test.x, test.y, b, test.compare < 0)
			// t.Errorf("Less(%v, %v) == %t, want %t", test.x, test.y, b, test.compare < 0)
		}
	}

	// Compare 函数比较两个类型为 T 的值 x 和 y 的大小关系，返回一个 整数 表示它们的比较结果。

	// 如果 x 和 y 都是 NaN，则返回 0；
	// 如果 x 是 NaN 或者 x 小于 y，则返回 -1；
	// 如果 y 是 NaN 或者 x 大于y，则返回 1；
	// 否则，返回 0 表示 x 等于 y。
	
	// from TestCompare
	for _, test := range tests {
		var c int
		switch test.x.(type) {
		case int:
			c = cmp.Compare(test.x.(int), test.y.(int))
		case string:
			c = cmp.Compare(test.x.(string), test.y.(string))
		case float64:
			c = cmp.Compare(test.x.(float64), test.y.(float64))
		case uintptr:
			c = cmp.Compare(test.x.(uintptr), test.y.(uintptr))
		}
		if c != test.compare {
			fmt.Printf("Compare(%v, %v) == %d, want %d\n", test.x, test.y, c, test.compare)
		}
	}

	// Or 函数，它接受一个或多个可比较的类型T的参数vals，并返回其中第一个不等于 零值 的参数。
	// 如果所有参数都是零值，那么它返回零值。

	// 1: 使用整数
	// fmt.Println(cmp.Or(0, 0, 5, 10)) 			// 输出：5

	// // 2: 使用字符串
	// fmt.Println(cmp.Or("", "hello", "world")) 	// 输出："hello"

	// // 3: 使用布尔值
	// fmt.Println(cmp.Or(false, true, false)) 	// 输出：true

	// // 4: 所有参数为零值
	// fmt.Println(cmp.Or(0, 0, 0)) 				// 输出：0 	(对于int类型，零值是0)
	// fmt.Println(cmp.Or(""))      				// 输出："" 	(对于string类型，零值是空字符串)
	// fmt.Println(cmp.Or(false))   				// 输出：false (对于bool类型，零值是false)

	ExampleOr()
	ExampleOr_sort()

	// from TestOr
	cases := []struct {
		in   []int
		want int
	}{
		{nil, 0},
		{[]int{0}, 0},
		{[]int{1}, 1},
		{[]int{0, 2}, 2},
		{[]int{3, 0}, 3},
		{[]int{4, 5}, 4},
		{[]int{0, 6, 7}, 6},
	}
	for _, tc := range cases {
		if got := cmp.Or(tc.in...); got != tc.want {
			fmt.Printf("cmp.Or(%v) = %v; want %v\n", tc.in, got, tc.want)
		}
	}
}

func ExampleOr() {
	// Suppose we have some user input
	// that may or may not be an empty string
	userInput1 := ""
	userInput2 := "some text"

	fmt.Println(cmp.Or(userInput1, "default"))
	fmt.Println(cmp.Or(userInput2, "default"))
	fmt.Println(cmp.Or(userInput1, userInput2, "default"))
	// Output:
	// default
	// some text
	// some text
}

func ExampleOr_sort() {
	type Order struct {
		Product  string
		Customer string
		Price    float64
	}
	orders := []Order{
		{"foo", "alice", 1.00},
		{"bar", "bob", 3.00},
		{"baz", "carol", 4.00},
		{"foo", "alice", 2.00},
		{"bar", "carol", 1.00},
		{"foo", "bob", 4.00},
	}
	// Sort by customer first, product second, and last by higher price
	slices.SortFunc(orders, func(a, b Order) int {
		return cmp.Or(
			cmp.Compare(a.Customer, b.Customer),
			cmp.Compare(a.Product, b.Product),
			cmp.Compare(b.Price, a.Price),
		)
	})
	for _, order := range orders {
		fmt.Printf("%s %s %.2f\n", order.Product, order.Customer, order.Price)
	}

	// Output:
	// foo alice 2.00
	// foo alice 1.00
	// bar bob 3.00
	// foo bob 4.00
	// bar carol 1.00
	// baz carol 4.00
}
