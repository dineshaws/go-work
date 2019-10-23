package main

import (
	"fmt"
)

func main() {
	a := 2
	fmt.Printf("decimal=%d, binary=%b and hexadecimal=%#x \n", a, a, a)
	a = a << 1
	fmt.Printf("decimal=%d, binary=%b and hexadecimal=%#x \n", a, a, a)
}
