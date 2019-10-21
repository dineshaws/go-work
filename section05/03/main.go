package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("value of x is=%v and value of y is %v and value of z is %v ", x, y, z)
	fmt.Println(s)
}
