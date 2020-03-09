package main

import (
	"fmt"
	"math"
)

func main() {
	su := square{
		length: 11.22,
	}
	ci := circle{
		radius: 15.35,
	}

	info(su)
	info(ci)
}

type square struct {
	length float64
}
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())

}
