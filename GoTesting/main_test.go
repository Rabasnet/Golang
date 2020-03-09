package main

import "testing"

// func main() {
// 	testCalculate()
//}
func testCalculate(t *testing.T) {
	if calculate(3) != 6 {
		t.Error(" the result should be 6")
	}
}

func anotherTest(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{4, 12},
		{5, 15},
		{6, 18},
	}
	for _, test := range tests {
		if output := calculate(test.input); output != test.expected {
			t.Error("the result should be ", test.expected)
		}
	}
}
