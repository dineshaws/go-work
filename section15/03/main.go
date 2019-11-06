package main

import (
	"fmt"
)

func main() {
	defer foo()
	bar()
}
func foo() {
	defer fmt.Println("foo called")
	defer f1()
	f2()
}
func bar() {
	fmt.Println("bar called")
}
func f1() {
	fmt.Println("f1 called")
}
func f2() {
	fmt.Println("f2 called")
}
