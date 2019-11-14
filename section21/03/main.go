package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Race condition example")
	counter := 0
	var wg sync.WaitGroup
	wg.Add(100)
	fmt.Println("GO routines ", runtime.NumGoroutine())
	for i := 0; i < 100; i++ {
		go func() {
			temp := counter
			temp++
			runtime.Gosched()
			counter = temp
			wg.Done()
		}()
		fmt.Println("GO routines ", runtime.NumGoroutine())
	}
	fmt.Println("GO routines ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("GO routines ", runtime.NumGoroutine())
	fmt.Println("FInal value of counter ", counter)
}
