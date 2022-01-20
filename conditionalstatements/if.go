package conditional

import "fmt"

func IfStatement() {
	number := 44
	booleanvalue := number == 75     //prints boolean value if it is true or false
	fmt.Println(booleanvalue)

	number1 := 44
	booleanvalue1 := number1 == 44   //prints boolean value if it is true or false
	fmt.Println(booleanvalue1)

	// #1
	if number%2 != 0 {        
	fmt.Println("yes")
	} else  {
	fmt.Println("no")
	}
	fmt.Println("out")

	//#2
	if true {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	//#3
	string := "statements"
	if string == "sample"{
		fmt.Println("1")
	} else {
		fmt.Println("2")
	}
	
	//Elseif

	sam := 3211
	if sam >= 4222 {                      /*if statement can have any number of elseif statements.Else statement is not mandatory */
		fmt.Println("greater")
	} else if sam < 2322 {
		fmt.Println("smaller")
	}else if sam >3211 {
	fmt.Println("greater and equal")     /*if none of the statements is true it prints nothing */
	}else {
		fmt.Println("done")
	}

	//Nestedif
	q := 33
	if q== 23 {
		if q==33 {
		fmt.Println("1") 
	}else {
	fmt.Println("2") 
	}
	}else if q >=33 {
		if q ==33 {
			fmt.Println("3")
		} else {
			fmt.Println("4")
		}
	}
}



