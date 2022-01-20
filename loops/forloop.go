package looping

import "fmt"

func ForLoop() {
	
	for i:=0 ; i <5 ; i++ {
		fmt.Println("Values of i :",i)
	}
	for i:=27 ; i <=37 ; i=i+5{                //loopinitializing,loop condition ,loop increment
			fmt.Println("Values of i :",i)
	}

	//Using Slices

	price := make([]int,0,21)
	for i := 0 ; i<=20 ; i=i+4 {
		price = append(price,i)
	}
	fmt.Println("price",price)
	fmt.Println("price of 2nd element",price[2])

	price1 := price[3]
	fmt.Println("price1",price1)
	
	

	
}