package main

import (
	//"github.com/prasannayanamadala/go-jan/go-january/arrays"
	//"github.com/prasannayanamadala/go-jan/go-january/practice"
	//"github.com/prasannayanamadala/go-jan/go-january/maps"
	//"github.com/prasannayanamadala/go-jan/go-january/conditionalstatements"
	//"github.com/prasannayanamadala/go-jan/go-january/loops"
	//"github.com/prasannayanamadala/go-jan/go-january/structures"
	"fmt"

	"github.com/prasannayanamadala/go-jan/go-january/functions"
)

func main() {
	
	//calculator.CalOperation()
	//array.ArrayDefine()
	//array.SliceDefine()
	//maps.MapsDefine()
	//conditional.IfStatement()
	//conditional.SwitchCondition()
	//looping.ForLoop()
	//structures.StructPractice()
	//structures.EmployeeDefine()
	a := 10
	b := 20
	result1 := functions_test.AddValues(a,b)
	fmt.Println(result1)

	a = 100
	b = 200
	result2 := functions_test.AddValues(a,b)
	fmt.Println(result2)
}

