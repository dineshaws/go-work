package main

import (
	"fmt"
)

func main() {
	b := map[string][]string{
		"gifts":   []string{"flower", "jhumka"},
		"rewards": []string{"files", "recharge"},
	}
	fmt.Println(b)
	b["bonus"] = []string{"holi", "diwali"}
	fmt.Println(b)
	delete(b, "rewards")
	fmt.Println(b)
	for key, record := range b {
		fmt.Println("this is for key=", key)
		for i, val := range record {
			fmt.Println("\t ", i, val)
		}
	}
}
