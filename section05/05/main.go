package main

import (
	"fmt"
)

type myInt int // int is underlying type

var x myInt // myInt is custom type
var y int

func main() {
	fmt.Printf("value of x is %v \t and type is %T  \n", x, x)
	x = 42
	fmt.Println(x)
	fmt.Printf("value of y is %v and type of y is %T \n", y, y)
	y = int(x)
	fmt.Printf("value of y is %v and type of y is %T\n", y, y)

}
