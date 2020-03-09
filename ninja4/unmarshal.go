package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//passing the values in the json format
	p := `[{"Name":"rabi basnet","EmployeeId":960},
		{"Name":"pyari chari","EmployeeId":969}]`
	//saving the values in the bytespace
	bs := []byte(p)
	//allocating the values of person struct into the people array
	var people []person

	//unmarshaling functions
	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}
	//print the values in the golang format
	fmt.Println("json to golang format\n", people)
}

type person struct {
	Name       string `json:"Name"`
	EmployeeID int    `json:"EmployeeId"`
}
