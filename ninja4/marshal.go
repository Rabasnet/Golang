package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	p1 := person{
		Name:       "rabi basnet",
		EmployeeId: 960,
	}
	p2 := person{
		Name:       "pyari chari",
		EmployeeId: 969,
	}
	p := []person{p1, p2}
	fmt.Println(p) //this is regular go lang printing
	//notice the fields Name and EmployeeId starts with capital letters
	//this is setup for marshal which is conversion from go to json
	bs, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(bs))
	a:=string(bs)
	fmt.Println("this is json format of the input data",a)
}

type person struct {
	Name       string
	EmployeeId int
}