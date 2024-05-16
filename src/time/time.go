package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	time.Sleep(time.Second)

	t := time.Now()

	elapsed := t.Sub(start)

	println(elapsed)			// 1005583000
	println(elapsed.Seconds())	// +1.005583e+000
	fmt.Printf("%.3f\n", elapsed.Seconds()) // 1.004
	println(int(elapsed.Seconds()))			// 1
	println(t.Date())    					// 2024511
	println(t.Format("2006-01-02"))			// 2024-05-11
	println(t.Format("2006-01-02 15:04:05"))// 2024-05-11 19:52:51
}