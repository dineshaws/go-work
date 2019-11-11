package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name    string   `json:"nm"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

func main() {
	fmt.Println("Unmarshalling")
	var users []user
	var jsonBlob = []byte(`[
        {"nm": "Dinesh", "age": 29, "hobbies": ["Cricket", "Reading"]},
        {"nm": "Prajapati", "age": 27, "hobbies": ["TV", "Sports"]}
    ]`)
	err := json.Unmarshal(jsonBlob, &users)
	if err != nil {
		fmt.Println("Error while Unmarshalling", err)
	}
	fmt.Println(users)
	for i, user := range users {
		fmt.Println("User #", i)
		fmt.Println("\n ", user.Name, user.Age)
	}
}
