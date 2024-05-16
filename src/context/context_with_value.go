package main

import (
	"context"
	"fmt"
)

type requestIDKey struct{}

func main() {
	/**
	// 应用示例，
	// 3、数据传递：
	// 在这里，context 用于在处理请求时携带请求ID，确保在整个请求处理过程中可用。 
	 */
	ctx := context.WithValue(context.Background(), requestIDKey{}, "12345")
	processRequest(ctx)
}

func processRequest(ctx context.Context) {
	reqID := ctx.Value(requestIDKey{}).(string)
	fmt.Printf("Processing request with ID: %s\n", reqID)
}