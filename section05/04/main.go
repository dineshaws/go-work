package main

import (
	"fmt"
)

type myInt int

var x myInt

func main() {
	fmt.Printf("value of x is %v \t and type is %T  \n", x, x)
	x = 42
	fmt.Println(x)

}
