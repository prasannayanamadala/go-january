package array

import "fmt"

func ArrayDefine() {

	 //general way of defining and declaring a variable
	grade1 := 30                   
	grade2 := 50
	grade3 := 70

	fmt.Println("grade1:",grade1)
	fmt.Println("grade2:",grade2)
	fmt.Println("grade3:",grade3)
	fmt.Println("grade1 address:",&grade1)
	fmt.Println("grade2 address:",&grade2)
	fmt.Println("grade3 address:",&grade3) //& gives address of a particular variable

	//firstway:shorthand declaration in an array

	grades := [3]int{grade1, grade2, grade3} //grades := [...]int{grade1, grade2, grade3}
	fmt.Println("grades in an array:",grades)
	fmt.Println("grade1 address:", &grades[0])
	fmt.Println("grade2 address:", &grades[1])
	fmt.Println("grade3 address:", &grades[2])
	
	firstgrade := grades[0]    //for accessing one value.Arrays are zero indexed
	fmt.Println("first grade representation in an array:",firstgrade)
	

	//second way:define

	var arraygit [7]string

	fmt.Println("string array:",arraygit) /*before declaring it will return an empty array */
	        
	//declare

	 arraygit[0] = "add"
	 arraygit[1] = "commit"
	 arraygit[2] = "push"
	 arraygit[3] = "pull"

	 fmt.Println("string array:",arraygit)
	 
	 /*shorthand using string array

		arraygit1 := "add"
		arraygit2 := "commit"
		arraygit3 := "push"
		arraygit4 := "pull"

		arraygithub := [5]string{arraygit1,arraygit2,arraygit3,arraygit4}
		fmt.Println("string array:",arraygithub) */

		//arrays can be created only if all elements have same datatype

arraygit1 := arraygit
arraygit[0] = "sub" //changing the value of one array 
fmt.Println("string array:",arraygit1)
fmt.Println("string array:",arraygit)
fmt.Println("string array length:",len(arraygit))

}