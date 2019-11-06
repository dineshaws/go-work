package main

import (
	"fmt"
)

func main() {
	fmt.Println("anonymous function example")
	f := func() int {
		return 2
	}
	fmt.Println(f())
}
