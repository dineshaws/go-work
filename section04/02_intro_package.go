package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello this is Println function use")
	n, err := fmt.Println("hello this is return values of function")
	fmt.Println(n, err)
	b, _ := fmt.Println("throw away err variable with void now no use of err wont throw error")
	fmt.Println(b)
}
