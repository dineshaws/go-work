package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		remainder := i % 4
		fmt.Printf("i=%d \t remainder=%v \n", i, remainder)
	}
}
