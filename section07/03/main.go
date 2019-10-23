// WAP to create TYPED and UNTYPED constants and print their values

package main

import "fmt"

const a = 12
const b string = "dinesh"
const (
	c     = true
	d int = 12
)

var (
	e int    = 32
	f string = "fas"
)

func main() {
	fmt.Println(a, b, c, d, e, f)
}
