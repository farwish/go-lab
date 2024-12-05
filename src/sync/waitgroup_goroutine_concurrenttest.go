package main

import (
	"sync"
	"testing"
)

func TestConcurrentAccess(t *testing.T) {
    var counter int
    var mu sync.Mutex
    var wg sync.WaitGroup

    numGoroutines := 100
    for i := 0; i < numGoroutines; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mu.Lock()
            counter++
            mu.Unlock()
        }()
    }

    wg.Wait()
    if counter != numGoroutines {
        t.Errorf("Expected counter to be %d, got %d", numGoroutines, counter)
    }
}