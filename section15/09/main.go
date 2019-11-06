package main

import (
	"fmt"
)

func main() {
	fmt.Println("callback function assignment")
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := sum(a...) // unfurling a slice
	fmt.Println("sum of all numbers ", s1)
	s2 := evenSum(sum, a)
	fmt.Println("sum of even numbers ", s2)
}

func evenSum(f func(x ...int) int, xi []int) int {
	var e []int
	for _, v := range xi {
		if v%2 == 0 {
			e = append(e, v)
		}
	}
	return f(e...)
}

func sum(xi ...int) int {
	var total int
	for _, v := range xi {
		total += v
	}
	return total
}
