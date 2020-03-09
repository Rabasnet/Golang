package main

import (
	"fmt"
	//"sync"
	"time"
)

//var wg sync.WaitGroup

func main() {
	//wg.Add(1)
	c := make(chan int)

	go coco(c)
	for msg := range c {
		fmt.Println(msg)
	}

	//wg.Wait()
}
func coco(c chan int) {

	for i := 1; i <= 5; i++ {

		c <- i * 5
		time.Sleep(time.Millisecond * 50)

	}
	//wg.Done()
	close(c) //sender function is closed so that the reciver side knows no more values are coming . Never close the reciver side
	//this can help avoid the deadlock
	//also, defining the buffer channel can help avoid the deadlock situation

}
