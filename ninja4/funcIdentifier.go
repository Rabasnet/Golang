package main

import (
	"fmt"
)

func main() {
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	a := foo(b...)
	fmt.Println(a)
	x := []int{10, 20, 22, 32, 35, 36, 65, 48, 41}
	y := bar(x)
	fmt.Println(y)
}
func foo(a ...int) int {
	sum := 0
	for _, value := range a {

		sum += value
		//fmt.Println(sum)
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, value := range x {
		sum += value

	}
	return sum
}
