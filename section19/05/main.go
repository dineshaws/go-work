package main

import (
	"fmt"
	"sort"
)

type user struct {
	Name string
	Age  int
}

// ByAge implements sort.Interface for []user based on
// the Age field.
type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	fmt.Println("Sort by Age and Name")
	u1 := user{
		Name: "Me",
		Age:  29,
	}
	u2 := user{
		Name: "My",
		Age:  25,
	}
	u3 := user{
		Name: "Myself",
		Age:  33,
	}
	users := []user{u1, u2, u3}
	fmt.Println(users)
	sort.Sort(ByAge(users))
	fmt.Println(users)
}
