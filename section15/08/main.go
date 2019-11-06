package main

import (
	"fmt"
)

func main() {
	fmt.Println("returning a function demo")
	f := foo()
	fmt.Println(f())
	fmt.Println(foo()())
}

func foo() func() int {
	return func() int {
		return 42
	}
}
