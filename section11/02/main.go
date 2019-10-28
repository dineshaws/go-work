package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	for i, v := range x {
		fmt.Printf("value %v on index %v\n", i, v)
	}
	fmt.Printf("type of composite literal slice %T\n", x)
}
