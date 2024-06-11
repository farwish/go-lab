package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
	$GOROOT/src/sync/oncefunc.go 
	
	OnceFunc, OnceValue, 和 OnceValues:
	高级封装（相比 once.Do(f)），提供了更方便的 API 来处理函数的单次执行，特别是当涉及到 返回值 或 panic 处理时。
	自动管理了 sync.Once 实例，无需用户直接操作 Once 变量。
	对于有返回值的情况（如 OnceValue 和 OnceValues），它们处理了存储和返回首次调用的结果。
	特别地，它们都妥善处理了函数 fpanic 的情况，确保在并发调用时 panic 行为一致且可预测，要么首次调用时立即重新 panic，要么在后续调用时再次 panic 同样的值。
	*/

	// 【1. OnceFunc】

	// 声明一个需要确保只执行一次的函数
	initFunc := func() {
		fmt.Println("初始化操作，这只应该打印一次!")
	}

	// 使用 OnceFunc 包装 initFunc
	wrappedInit := sync.OnceFunc(initFunc)

	// 启动多个goroutine同时调用 wrappedInit
	for i := 0; i < 10; i++ {
		go func() {
			wrappedInit()
		}()
	}

	// 等待一段时间，确保所有goroutine有机会执行
	time.Sleep(1 * time.Second)
	fmt.Println("=====================================/")

	// 【2. OnceValue】

	// 声明一个需要确保只执行一次并返回值的函数
	getDatabaseConnection := func() string {
		fmt.Println("建立数据库连接，这只应该执行一次!")
		return "(数据库连接字符串)"
	}

	// 使用 OnceValue 包装 getDatabaseConnection
	getConnOnce := sync.OnceValue(getDatabaseConnection)

	// 启动多个 goroutine 同时尝试获取数据库连接字符串
	var connStrings []string
	for i := 0; i < 10; i++ {
		go func() {
			conn := getConnOnce()
			connStrings = append(connStrings, conn)
		}()
	}

	// 等待一段时间，确保所有 goroutine 有机会执行
	time.Sleep(1 * time.Second)

	// 打印所有获取到的数据库连接字符串，应该全部相同
	fmt.Println("所有获取到的数据库连接字符串:", connStrings)
	fmt.Println("所有可能的并发调用已完成。")
	fmt.Println("=====================================/")


	// 【3. OnceValues】

	// 声明一个需要确保只执行一次并返回两个值的函数
	loadConfig := func() (string, int) {
		fmt.Println("加载配置文件，这只应该执行一次!")
		return "配置数据A", 123 // 假设加载的配置为一个字符串和一个整数
	}

	// 使用 OnceValues 包装 loadConfig
	loadConfigOnce := sync.OnceValues(loadConfig)

	// 启动多个 goroutine 同时尝试加载配置
	var configs []struct {
		strVal string
		intVal int
	}
	for i := 0; i < 10; i++ {
		go func() {
			configA, configB := loadConfigOnce()
			configs = append(configs, struct {
				strVal string
				intVal int
			}{configA, configB})
		}()
	}

	// 等待一段时间，确保所有 goroutine 有机会执行
	time.Sleep(1 * time.Second)

	// 打印所有获取到的配置信息，应该全部相同
	fmt.Println("所有获取到的配置信息:")
	for _, config := range configs {
		fmt.Printf("字符串值: %s, 整数值: %d\n", config.strVal, config.intVal)
	}
	fmt.Println("所有可能的并发调用已完成。")
}