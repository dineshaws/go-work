package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Mutax example")
	count := 100
	counter := 0
	fmt.Println("Goroutines", runtime.NumGoroutine())
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			mu.Lock()
			temp := counter
			temp++
			//runtime.Gosched()
			counter = temp
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Goroutines", runtime.NumGoroutine())
	fmt.Println("final counter value ", counter)
}
