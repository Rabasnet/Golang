package main

import "fmt"

func main() {
	m := map[string][]string{
		`rabi`:   []string{`car`, `money`, `moreMoney`},
		`hem`:    []string{`samsung`, `iphone`, `hp`},
		`ummesh`: []string{`macbook`, `googlePixel`, `earphone`},
	}
	m[`torilahure`] = []string{`kutta`, `kamina`, `aaulad`}
	delete(m, `rabi`)
	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Printf("index: %v \t\t value: %v\n", i, v2)
		}

	}

}

// var n:=m[`torilahure`]:    []string{`kutta`, `kamina`, `aaulad`}
// for l, vl := range n {
// 	fmt.Println(k)
// 	for i, v3 := range vl {
// 		fmt.Printf("index: %v \t\t value: %v\n", i, v3)
// 	}

// }

// }
