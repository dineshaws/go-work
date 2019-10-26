package main

import (
	"fmt"
)

func main() {
	counter := 1
	for i := 65; i <= 90; i++ {
		fmt.Println(counter)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t %#U \n", i)
		}
		counter++
	}
}
