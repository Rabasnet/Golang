package main
import (
	"fmt"
	"runtime"
	"sync"

)



func main(){
	counter:=0
	const gs =10
	var w sync.WaitGroup
	var m sync.Mutex
	for i:=0 ; i<gs; i++{
		w.Add(gs)
		
		go func (){
			m.Lock()
			v:=counter
		fmt.Println("No. of goroutines" , runtime.NumGoroutine())
		//runtime.Gosched()
		v++
		counter=v
		m.Unlock()
		w.Done()
		
		}()
	}
	
	w.Wait()
	fmt.Println("goroutines", runtime.NumGoroutine)
	fmt.Println("count", counter)
}

