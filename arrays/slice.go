package array

import "fmt"

func SliceDefine() {

	slices := make([]int,0)
	
	slices = append(slices,10,20,30,40,50,60,70,80,90,100)
	fmt.Println(slices)   //0, 1, 2, 3, 4, 5
	
	//slices = slices[2:6] //second element to fifth(6-1) element
	//fmt.Println(slices)

	slices = append(slices,110,120,130,140)
	fmt.Println(slices)
	fmt.Println(len(slices))
	fmt.Println(&(slices[10]))

	slices = append(slices,150,160,170,180)
	fmt.Println(slices)
	}