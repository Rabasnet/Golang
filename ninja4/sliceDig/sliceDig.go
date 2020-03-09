package main

import (
	"fmt"
)

func main() {
	//4 ways to create a slice
	//1. initialize with variables
	r := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//2.initialize without declaring the variable
	a := []int{}
	//3.decalre but don't allocate memory
	var b []int
	//4.initialize with no value, allocate memory, length and capacity
	i := make([]int, 5, 10)

	//printing a portion of the slice
	fmt.Println(r[2:9])

	//delete a portion
	m := copy(i, r[2:8]) // will be copied to slice i, the length of i shouuld be defined first
	fmt.Println(m)       //this gives the number of items copied
	fmt.Println(i)

	i[2] = 10 //changing the value
	fmt.Println(i)

	//appending in the slice
	a = append(a, 12, 13, 14, 15, 16)
	fmt.Println(a)
	//subslicing
	a1 := a[0:3]
	fmt.Println(a1)
	a1[0] = 22 //subslicing changes the value in original slice too
	fmt.Println(a1)
	fmt.Println(a)

	//deleting element from slice
	fmt.Println(r)
	r = append(r[0:2], r[3:9]...) //variadiac//need to put the slice name in the parameter
	fmt.Println(r)
	//appending two slices//concatinate
	r = append(r, a...)
	fmt.Println(r)
	b = append(r, a...)
	fmt.Println(b) //this gives modified r with a  and plus a

}
