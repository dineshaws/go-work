package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First string
	Age   int
}

func main() {
	fmt.Println("Marshalling")
	u1 := user{
		First: "Dinesh",
		Age:   29,
	}
	u2 := user{
		First: "Me",
		Age:   25,
	}
	users := []user{u1, u2}
	fmt.Println(users)
	ba, err := json.Marshal(users)

	if err != nil {
		fmt.Println("Error in Marshalling ", err)
	}
	fmt.Println(string(ba))
	os.Stdout.Write(ba)
}
