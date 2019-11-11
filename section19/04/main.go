package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Sort integers and strings")
	xi := []int{2, 4, 1, 4, 6, 7, 4, 3, 2, 6, 8, 0, 34, 23, 11, 43}
	fmt.Println("unsorted xi ", xi)
	sort.Ints(xi)
	fmt.Println("sorted xi ", xi)
	xs := []string{"Kumar", "Kumawat", "Prajapati", "Dinesh"}
	fmt.Println("unsorted xs ", xs)
	sort.Strings(xs)
	fmt.Println("sorted xs ", xs)

}
