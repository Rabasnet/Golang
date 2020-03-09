package main

import (
	"fmt"
)

func main() {

	a := [][]string{{"james", "Bond", "Shaken", "not stirred"}, {"miss", "moneypenny", "Helllo", "James"}}
	for _, b := range a {
		fmt.Println(b)
		for i, c := range b {
			fmt.Printf("index position: %v\t\t index value: %v\n", i, c)
		}

	}
}
