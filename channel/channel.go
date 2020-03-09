//INTRODUCTION TO GO , PART 15 : CHANNELS (youtube video)
//YOU SHOULD ALWAYS HAVE A STRATEGY FOR HOW YOUR GO ROUTINE IS GOING TO SHUT DOWN WHEN YOU CREATE A GO ROUTINE  OTHERWISE IT CAN BE A SUBTLE RESOURCE LEAK AND EVENTUALLY THAT
//LEAK CAN BRING THE APPLICATION DOWN
// var doneCh:= make (chan struct{}{}) 2 empty braces;  here, empty struct saves couple of memory allocations

//avoid the race conditions and memory sharing problem
//pass data between the go routines and restrict the data flow
//here we are doing the multi-direction channel, uni directional channel, buffered channel, closing the channel,
//for ...range loop and the select statements(like switch statemets but designed for channels)
//2 types of synchroziation here, waitgroup synchrozination waits for the main go routine to finish and another channel synchrozination to manage th data flow between them
//dedicate a go routine to either reading from channel or writing from a channel, don't make a go routine do both the task
package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 20)

	wg.Add(2)
	go func(ch <-chan int) { //reciver only channel go routine
		for i := range ch { //instead of the for loop, we can use the comma ok idiom too which checks the conditions manuallly BUT the for loop will detect the channel closure automatically
			fmt.Println(i) // use comma OK idiom when the reciver contain spinning off of the go routines and the use of the range may not make sense
		}

		// i = <-ch // here i is already declared and not new variable so no need to declare
		// fmt.Println(i)
		//instead of retriving the data one by one, we can loop it with for loop

		wg.Done()
	}(ch)
	//  you have to pass the channel as argumets //here you are passing bidirectional channel to the uni directional channel. this is specific to the channel only

	go func(ch chan<- int) { //sender only channel go routine
		ch <- 42
		ch <- 44
		close(ch) //if we dont close the channel, it will cause deadlofk because the for loop in the reciver function keeps on expecting the values and causes the deadkock
		//when you close the channel, nth else is send after it. If you try to send a message after closing the channel, it will give the PANIC signal
		//in that case, you can have a deferred function with a recovery.....

		wg.Done()

	}(ch)

	wg.Wait()

}
