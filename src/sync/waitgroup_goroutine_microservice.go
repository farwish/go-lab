package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func serviceHandler(w http.ResponseWriter, r *http.Request) {
    ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
    defer cancel()

    go longRunningTask(ctx, w)
}

func longRunningTask(ctx context.Context, w http.ResponseWriter) {
    select {
    case <-ctx.Done():
        fmt.Fprintf(w, "Request timed out")
    case <-time.After(5 * time.Second):
        fmt.Fprintf(w, "Long running task completed")
    }
}

func main() {
    http.HandleFunc("/", serviceHandler)
    http.ListenAndServe(":8080", nil)
}