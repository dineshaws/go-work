package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, " ", p.last, " age is ", p.age)
}
func main() {
	p := person{
		first: "Dinesh",
		last:  "Prajapat",
		age:   28,
	}
	p.speak()
}
