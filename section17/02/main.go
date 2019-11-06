package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	fmt.Println("Change value at address using pointers")
	p1 := person{
		first: "Dinesh",
		last:  "Prajaapti",
		age:   29,
	}
	fmt.Println("before change p", p1)
	changeMe(&p1)
	fmt.Println("after change p", p1)

}

func changeMe(p *person) {
	*p = person{
		first: "Change",
		last:  "Name",
		age:   24,
	}
	fmt.Println("change me after p", *p)
	(*p).first = "Dereference"
	p.last = "Selector" // can use p.last instead of (*p).last, it also works
}
