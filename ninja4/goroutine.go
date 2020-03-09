package main

import (
	"fmt"
	"runtime"
	"sync"
)

var w sync.WaitGroup

func main() {
	fmt.Println("no of cpu\t\t", runtime.NumCPU())
	fmt.Println("no of go routines\t", runtime.NumGoroutine())
	w.Add(2)
	go ghi()

	go def()
	abc()
	fmt.Println("no of cpu\t\t", runtime.NumCPU())
	fmt.Println("no of go routines\t", runtime.NumGoroutine())
	w.Wait()
}

func abc() {
	value := 0
	for i := 0; i < 5; i++ {
		value += i
		fmt.Printf("the index no. is %v and the total value is %v\n", i, value)
	}

}

func def() {
	value := 0
	for i := 6; i <= 10; i++ {
		value += i
		fmt.Printf("the index no. is %v and the total value is %v\n", i, value)

	}
	w.Done() //this done func  should be just before wrapping the go function i.e before closing the function
}
func ghi() {
	value := 0
	for i := 11; i <= 15; i++ {
		value += i
		fmt.Printf("the index no. is %v and the total value is %v\n", i, value)

	}
	w.Done() //this done func  should be just before wrapping the go function i.e before closing the function
}
