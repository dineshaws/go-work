package main

import (
	"fmt"
)

func main() {
	switch false {
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	default:
		fmt.Println("default")
	}
}
