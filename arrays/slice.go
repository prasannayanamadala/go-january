package array

import "fmt"

func SliceDefine() {

	slices := make([]int,0,20)
	
	slices = append(slices,10,20,30,40,50,60,70,80,90,100)
	fmt.Println(slices)   //0, 1, 2, 3, 4, 5
	
	slices = slices[2:6] //second element to fifth(6-1) element
	fmt.Println(slices)   //30,40,50,60

	slice1 := slices
	slices[2] = 333     //30,40,333,60
	fmt.Println("slice1:",slice1)
	fmt.Println(slices) //30,40,333,60

	slices = append(slices,110,120,130,140)
	fmt.Println(slices) //30,40,333,60,110,120,130,140
	
	fmt.Println("length of slices:",len(slices))  //8
	fmt.Println("capacity of slices:",cap(slices)) 

	slice2 := slices[2:]
	fmt.Println("slice2:",slice2) // 333,60,110,120,130,140

	slices = append(slices,150,160,170,180)
	fmt.Println(slices) //30,40,333,60,110,120,130,140,150,160,170,180

	slice3 := slices[:7]
	fmt.Println("slice3:",slice3) // 30,40,333,60,110,120,130

	slices[4] = 500
	fmt.Println(slices) //30,40,333,60,500,120,130,140,150,160,170,180
	
	}