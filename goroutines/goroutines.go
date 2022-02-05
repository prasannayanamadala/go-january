package goroutine

import (
	"fmt"
	"time"
)
		

func GoRoutineTest() {
	fmt.Println("testing of goroutine")
	go ParallelismTest()
	time.Sleep(2 * time.Millisecond)
}
func ParallelismTest() {
	time.Sleep(2 * time.Millisecond)
	fmt.Println("Does it print on time")
}