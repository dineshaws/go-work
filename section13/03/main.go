package main

import (
	"fmt"
)

type vehicle struct {
	color string
	doors int
}
type truck struct {
	fourWheel bool
	vehicle
}
type sedan struct {
	luxary bool
	vehicle
}

func main() {
	t := truck{
		fourWheel: true,
		vehicle: vehicle{
			color: "White",
			doors: 2,
		},
	}
	s := sedan{
		luxary: true,
		vehicle: vehicle{
			color: "Grey",
			doors: 4,
		},
	}
	fmt.Println(t)
	fmt.Println(t.fourWheel)
	fmt.Println(t.color)
	fmt.Println(t.doors)
	fmt.Println(s)
	fmt.Println(s.luxary)
	fmt.Println(s.color)
	fmt.Println(s.doors)

}
