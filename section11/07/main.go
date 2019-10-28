package main

import (
	"fmt"
)

func main() {
	xs1 := []string{"James", "Bond"}
	xs2 := []string{"Iron", "Man"}
	fmt.Println(xs1)
	fmt.Println(xs2)
	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)

	for i, record := range xxs {
		fmt.Println(i)
		for j, item := range record {
			fmt.Printf("\t%v.%v %v\n", i, j, item)
		}
	}
}
