package main

import (
	"fmt"
)

func main() {
	fmt.Println("Closure")
	a := incrementor()
	fmt.Println("a called ", a())
	fmt.Println("a called ", a())
	fmt.Println("a called ", a())
	fmt.Println("a called ", a())
	fmt.Println("a called ", a())
	b := incrementor()
	fmt.Println("b called ", b())
	fmt.Println("b called ", b())
	fmt.Println("b called ", b())
}

func incrementor() func() int {
	var x int = 12
	return func() int {
		x++
		return x
	}
}
