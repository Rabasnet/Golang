package main

import "fmt"

func calculate(x int) (result int) {
	result = x * 2
	return result

}
func main() {

	result := calculate(5)
	fmt.Println("the total is ", result)
}
