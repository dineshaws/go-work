package main

import (
	"fmt"
)

func main() {
	fmt.Println("Recursion")
	fmt.Println("factorial of number 4", factorial(4))
	fmt.Println("factorial of number 5", factorial(5))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
