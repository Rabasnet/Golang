package main

import (
	"fmt"
)

func main() {
	a := country{
		name:     "USA",
		size:     50,
		location: "North America",
	}
	a.America()
}

type country struct {
	name     string
	size     int
	location string
}

func (a country) America() {

	//b:=(a.name , a.size , a
	fmt.Println(a.name, a.size, a.location)
}
