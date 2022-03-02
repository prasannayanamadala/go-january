package soaps

import "fmt"

func SoapBars() {
	array:=[]int{2,5,10}

	smallestNumber := array[0]

	for _,element :=range array {
		if element < smallestNumber {
			smallestNumber = element

		}
	}
	fmt.Println("smallest number of bars:",smallestNumber)
}