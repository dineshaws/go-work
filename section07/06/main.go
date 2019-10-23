package main

import (
	"fmt"
)

const year = 2016
const (
	a uint32 = (iota + year)
	b uint32 = iota + year
	c uint32 = iota + year
	d uint32 = iota + year
)

func main() {
	fmt.Println(a, b, c, d)
}
