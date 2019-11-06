package main

import (
	"fmt"
)

func main() {
	fmt.Println(foo())
	x, y := bar()
	fmt.Println(x, y)
	fmt.Printf("Type of x %T\n", x)
	fmt.Printf("Type of y %T\n", y)
}

func foo() int {
	return 23
}

func bar() (int, string) {
	return 24, "bar"
}
