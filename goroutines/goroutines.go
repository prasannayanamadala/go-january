package goroutine

import (
	"fmt"
	//"time"
	"sync"
)
		
var wg = sync.WaitGroup{}

func GoRoutineTest() {
	fmt.Println("testing of goroutine")
	x := 1
	for i:=0; i<5 ; i++ {
	
	wg.Add(2)
	go ParallelismTest(&x)
	go NextFunc(&x)
	wg.Wait()
	}
	//time.Sleep(2 * time.Millisecond)
}
func ParallelismTest(a *int) {
	//time.Sleep(2 * time.Millisecond)
	*a = *a +1
	fmt.Println("1st func:",*a)
	wg.Done()
}
func NextFunc(a *int) {
	*a = *a +1
	fmt.Println("2nd func:",*a)
	wg.Done()
}