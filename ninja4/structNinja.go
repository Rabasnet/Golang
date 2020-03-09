package main

import (
	"fmt"
)

func main() {
	p := prson{
		first: "rabi",
		last:  "banset",
		age:   30,
	}
	speak(p)
}

type prson struct {
	first string
	last  string
	age   int
}

func speak(p prson) {

	fmt.Println("my name is", p.first, p.last, "and i am ", p.age, "years old")

}
