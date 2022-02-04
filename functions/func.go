package functions_test

import "fmt"

func AddValues(a int,b int)int{

//a := 10
//b := 20
c := a+b

fmt.Println(a)
return c
}

func AddThreeValues(f int , str string,z bool)(int,string,bool){

	//f = 500
	//str = "snow storm"
	//z = false
	//fmt.Println(f)
	//fmt.Println(str)
	//fmt.Println(z)
	return f,str,z
}
func DeferKeyword()int { //defer is a keyword that keeps the statement or that line executed in the last after executing the entire program
	fmt.Println(1)
	fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4) //if there are two defer statements in one program,the statements that are last in with execute first(last in first out)
	fmt.Println(5)
	fmt.Println(6)
	return 5
}
func PanicKeyword(m int,n int)int {
	defer Recover()
	var p int
	if n ==0 {
	panic("n value cannot be 0")
	} else {
	p =m/n
	}
	fmt.Println("this is the way")
	return p
}
func Recover() {
	if r := recover(); r != nil {
		fmt.Println("Recovered")
	}
}


