package goroutine

import (
	"fmt"
	//"time"
	"sync"
)
		
var wg = sync.WaitGroup{}

func GoRoutineTest() {
	fmt.Println("testing of goroutine")
	wg.Add(2)
	go ParallelismTest()
	go NextFunc()
	wg.Wait()
	//time.Sleep(2 * time.Millisecond)
}
func ParallelismTest() {
	//time.Sleep(2 * time.Millisecond)
	fmt.Println("Does it print on time")
	wg.Done()
}
func NextFunc() {
	fmt.Println("think so,it can print on time")
	wg.Done()
}