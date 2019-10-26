package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 100 {
		i++
		r := i % 3
		if r == 1 {
			fmt.Printf("When %v divided by 3, remainder aka modulus is 1\n", i)
		} else if r == 2 {
			fmt.Printf("When %v divided by 3, remainder aka modulus is 2\n", i)
		} else {
			fmt.Printf("When %v divided by 3, remainder aka modulus is 0\n", i)
		}
	}
}
