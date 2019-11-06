package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers")
	var a int = 23
	fmt.Println("value of a   ", a)
	fmt.Println("address of a ", &a)
	fmt.Println("value of a ", *&a)
}
