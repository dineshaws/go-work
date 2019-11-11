package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type user struct {
	Name    string
	Age     int
	Hobbies []string
}

func main() {
	fmt.Println("Writer interface")
	fmt.Fprintln(os.Stdout, "Writer interface")
	io.WriteString(os.Stdout, "Writer interface\n")

	u1 := user{
		Name:    "DInesh",
		Age:     29,
		Hobbies: []string{"Cricket", "Reading"},
	}
	u2 := user{
		Name:    "Prajapati",
		Age:     27,
		Hobbies: []string{"Sports"},
	}
	users := []user{u1, u2}
	fmt.Println(users)
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("Json encoder error ", err)
	}
}
