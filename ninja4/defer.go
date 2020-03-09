package main

import (
	"fmt"
)

func main() {
	defer rabi()
	defer rabi1()
	rabi2()

}

func rabi() {
	defer rabi3()
	fmt.Println("Hello, rabi")

}
func rabi1() {
	fmt.Println("Hello, rabi1")
}
func rabi2() {
	fmt.Println("Hello, rabi2")
}
func rabi3() {
	fmt.Println("Hello, rabi3")
}
