package main

import "fmt"

type person struct {
	fName        string
	lName        string
	icreamFlavor string
}

func main() {
	p1 := person{
		fName:        "rabi",
		lName:        "basnet",
		icreamFlavor: "vanilla",
	}
	p2 := person{
		fName:        "sandhya",
		lName:        "khwanju",
		icreamFlavor: "choclate",
	}
	// fmt.Println(p1)
	// fmt.Println(p2)
	a := map[string]person{
		p1.lName: p1,
		p2.lName: p2,
	}
	fmt.Println(a)
}
