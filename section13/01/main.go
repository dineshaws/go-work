package main

import (
	"fmt"
)

type person struct {
	first   string
	last    string
	flavors []string
}

func main() {

	p1 := person{
		first:   "Dinesh",
		last:    "Prajapati",
		flavors: []string{"chocolate", "vanilla"},
	}
	p2 := person{
		first:   "foo",
		last:    "bar",
		flavors: []string{"cashew"},
	}
	fmt.Println(p1.first, p1.last)
	for i, flavor := range p1.flavors {
		fmt.Println(i, flavor)
	}
	fmt.Println(p2.first, p2.last)
	for i, flavor := range p2.flavors {
		fmt.Println(i, flavor)
	}

}
