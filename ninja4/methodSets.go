package main

import "fmt"

func main() {
	p1 := person{
		name: "rabi basnet",
	}
	p1.speak()
}

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("Hello there")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
