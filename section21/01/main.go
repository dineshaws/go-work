package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Start time", time.Now())
	fmt.Println("GO routines")
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("ARCH: ", runtime.GOARCH)
	fmt.Println("CPUS: ", runtime.NumCPU())
	fmt.Println("Routines: ", runtime.NumGoroutine())

	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
	fmt.Println("Routines: ", runtime.NumGoroutine())
	fmt.Println("End time", time.Now())
}
func foo() {
	for i := 0; i < 20; i++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 20; i++ {
		fmt.Println("bar: ", i)
	}
	wg.Done()
}
