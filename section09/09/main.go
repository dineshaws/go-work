package main

import (
	"fmt"
)

func main() {
	var favSport string = "badminton"
	switch favSport {
	case "cricket":
		fmt.Println("cricket is fav sport")
	case "soccer":
		fmt.Println("soccer is fav sport")
	case "badminton":
		fmt.Println("badminton is fav sport")
	default:
		fmt.Println("not interested in sport")
	}
}
