package main

import (
	"fmt"
)

func main() {
	abc()
}
func abc() {
	val := []int{0, 2, 1, 5, 6, 7, 8}
	lambai := len(val)

	for i := 0; i <= lambai-1; i++ {
		for j := 0; j <= lambai-1; j++ {
			result := val[i] + val[j]
			if result == 9 {
				fmt.Println("result:", result, "values", val[i], val[j], "\t\t\t index position", i, j)
			}

		}

	}
}
