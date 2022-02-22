package goroutine

import (
	"fmt"
	//"time"
	"sync"
)
		
var wg = sync.WaitGroup{}

//func GoRoutineTest() {
	//fmt.Println("testing of goroutine")
	//x := 1

 	//for i:=0; i<5 ; i++ {
	
// 	wg.Add(2)
// 	go ParallelismTest(&x)
// 	go NextFunc(&x)
// 	wg.Wait()
// 	}
// 	//time.Sleep(2 * time.Millisecond)
// }
// func ParallelismTest(a *int) {
// 	//time.Sleep(2 * time.Millisecond)
// 	*a = *a +1
// 	fmt.Println("1st func:",*a)
// 	wg.Done()
// }
// func NextFunc(a *int) {
// 	*a = *a +1
// 	fmt.Println("2nd func:",*a)
// 	wg.Done()
// }
//ch := make(chan int,1)
//ch <-x
//wg.Add(1)
//go SampleGoroutine(ch)
//wg.Wait()
//}
//func SampleGoroutine(ch chan int) {
	//val := <-ch
	//fmt.Println(val)
	//wg.Done()
//}
func GoRoutineTest() {
	//fmt.Println("testing of goroutine")
	
	ch := make(chan int,2)
	x := 1
	for i:=0; i<5 ; i++ {
	ch <-x
 	wg.Add(2)
 	go ParallelismTest(ch)
	go NextFunc(ch)
 	wg.Wait()
 	}

 }
 func ParallelismTest(ch chan int) {

	val := <-ch
 	fmt.Println("1st func:",val)
 	wg.Done()
 }
 func NextFunc(ch chan int) {
 	val1 := <-ch
	fmt.Println("2nd func:",val1)
 	wg.Done()
 }
