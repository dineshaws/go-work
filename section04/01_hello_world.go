package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world, first program of the this section")
	foo()
	fmt.Println("sequential code")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Printf("here is the even number %d \n", i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("in foo function")
}
func bar() {
	fmt.Println("program exit here")
}

// control flow
// 1. sequential
// 2. Iteration
// 3. conditional
