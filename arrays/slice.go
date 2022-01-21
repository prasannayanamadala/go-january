package array

import "fmt"

func SliceDefine() {

	slices := make([]int,0,20)
	
	//slices = append(slices,10,20,30,40,50,60,70,80,90,100)
	//fmt.Println(slices)   //0, 1, 2, 3, 4, 5

	for i:=10 ; i<=100 ;i= i+10 {          //using for loop for declaring if values that are assigning in sequential order only
		slices = append(slices,i)
	}
	fmt.Println(slices)
	
	for i:= 0 ; i<=len(slices)-1 ; i=i+1 {
		if i==4 {
			slices[0]=100
		}
		fmt.Println(slices[i])
	}
	//fmt.Println(slices[i])  //i value can be printed only inside the for loop

	/* for range loop */                   //this kind of for loop is used in maps but not in slices 
		for index, val :=range slices {
		fmt.Println()
		fmt.Print("index:",index)
		fmt.Print("   ")
		fmt.Print("val:",val)
	}
	
	for _, val :=range slices {  //in for loop if only one value has to be printed use _(viceversa)
		fmt.Println()
		fmt.Print("   ")
		fmt.Print("val:",val)
	}
		slices = slices[2:6]     //second element to fifth(6-1) element
	fmt.Println(slices)       //30,40,50,60

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