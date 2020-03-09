package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	isFourWheel bool
}
type sedan struct {
	vehicle
	isLuxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		isFourWheel: true,
	}
	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		isLuxury: true,
	}
	fmt.Println(t1)
	fmt.Println(s1)
	fmt.Println(s1.vehicle.doors)
	fmt.Println(s1.vehicle.color)
	fmt.Println(s1.isLuxury)
	fmt.Println(t1.vehicle.doors)
	fmt.Println(t1.vehicle.color)
	fmt.Println(t1.isFourWheel)
}
