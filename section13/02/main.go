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
		first: "Dinesh",
		last:  "Prajapati",
		flavors: []string{
			"Vanilla",
			"Chocolate",
		},
	}
	p2 := person{
		first: "Raju",
		last:  "Kumar",
		flavors: []string{
			"Cashew",
			"Strawberry",
		},
	}
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	for key, value := range m {
		fmt.Println(key)
		fmt.Println(value)
	}
}
