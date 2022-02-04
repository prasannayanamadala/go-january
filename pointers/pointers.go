package pointers_test

import "fmt"

//func PointerDefine() {

	//var str string = "hello"
	//fmt.Println(str)
	//fmt.Println(&str)
	//var ptr *string = &str
	//fmt.Println(ptr)
	//fmt.Println(*ptr)
	//fmt.Println(*&str)

	//var a int =100
	//fmt.Println(a)
	//fmt.Println(&a)
	//var ptr1 *int = &a
	//fmt.Println(ptr1)
	//fmt.Println(*ptr1)
	//fmt.Println(*&a)

// 	a := 100
// 	b := 200
// 	Sample(a,b) //100,200
// 	fmt.Println(a) //100,200
// 	fmt.Println(b)
// 	sample_test(a,b) //100,200
// 	fmt.Println(a)
// 	fmt.Println(b)
// }
// 	func Sample(a int,b int) { //100,200
// 		a = a+100     //200,200
// 		fmt.Println(a)
// 		fmt.Println(b)
// 	}
// 	func sample_test(a int,b int) { //100,200
// 		a =a+2  //102,200
// 		fmt.Println(a)
// 		fmt.Println(b)

// 	}

// 	a := 100
// 	b := 200
// 	Sample(a,b) //100,200
// 	fmt.Println(a) //100,200
// 	fmt.Println(b)
// 	sample_test(a,b) //100,200
// 	fmt.Println(a)
// 	fmt.Println(b)
// }
// 	func Sample(a int,b int) { //100,200
// 		a = a+100     //200,200
// 		fmt.Println(a)
// 		fmt.Println(b)
// 	}
// 	func sample_test(a int,b int) { //100,200
// 		a =a+2  //102,200
// 		fmt.Println(a)
// 		fmt.Println(b)

//	}

// a := 100
// b := 200
// Sample(&a,&b) 
// fmt.Println(a) //200,200
// fmt.Println(b)
// sample_test(&a,&b) 
// fmt.Println(a)//202,200
// fmt.Println(b)
// }
// func Sample(a *int,b *int) { //100,200
// 	*a = *a +100     //200,200
// 	fmt.Println(*a)
// 	fmt.Println(*b)
// }
// func sample_test(a *int,b *int) { //200,200
// 	*a =*a+2  //202,200
// 	fmt.Println(*a)
// 	fmt.Println(*b)

//}                              //variables are pass by values
// func CheckPassby() {         //Arrays are pass by values

// 	var arr [4]int
// 	arr[0] = 4
// 	arr[1] = 1
// 	arr[2] = 2
// 	arr[3]= 3
// 	fmt.Println(arr) //[4 1 2 3]
// 	NextCheck(arr)
// 	fmt.Println(arr) //[4 1 2 3]
// }
// func NextCheck(arr[4] int) {
// 	arr[3] = 5      
// 	fmt.Println(arr)   //[4 1 2 5]
// } 
// }
// func SliceCheckifpassby(){

// 	slice := make([]int,0)
// 	slice = append(slice, 1,2,3,4,5)
// 	fmt.Println(slice)    //[1 2 3 4 5]
// 	NextSlice(slice)
// 	fmt.Println(slice)    //[1 2 6 4 5]     //slices are passby reference
// }
// func NextSlice(slice []int) {

// 	slice[2] = 6
// 	fmt.Println(slice)  //[1 2 6 4 5]     
// }
// func MapsCheckpassby() {
 
// 	simplemap := make(map[string]int)
// 	simplemap["A"] = 1
// 	simplemap["B"] = 2
// 	simplemap["C"] = 3
// 	simplemap["D"] = 4
// 	simplemap["E"] = 5
// 	simplemap["F"] = 6
// fmt.Println(simplemap) //[A:1 B:2 .....]
// 	MapCheck(simplemap)
// 	fmt.Println(simplemap) //[A:2 B:2 ....]      //maps are passby reference
// }
// func MapCheck (simplemap map[string]int){
// 	simplemap["A"] =2
// 	fmt.Println(simplemap)//[A:2 B:2 ...]


//}
type Employee struct {
	Name string
	Salary int
	Position string
	//Address EmployeeAddress
} 
// type EmployeeAddress struct {
// 	Housenumber int
// 	Street string
// 	City string
// 	State string
// 	Zipcode int
// }

 func StructCheckifpassby() {              //structures are passby value
	 employees := make([]Employee,0)
	
	var emp1 Employee
	emp1.Name = "jenny"
	emp1.Salary = 25000
	emp1.Position = "Designer"
 
	// var emp1Address EmployeeAddress
	// emp1Address.Housenumber = 638
	// emp1Address.Street = "woodbury ct"
	// emp1Address.City = "saltlake city"
	// emp1Address.State = "utah"
	// emp1Address.Zipcode = 56768

	// emp1.Address = emp1Address
	
	//employees = append(employees,emp1)
	//fmt.Println(employees)

	var emp2 Employee
	emp2.Name = "mary"
	emp2.Salary = 30000
	emp2.Position = "cashier"

	// var emp2Address EmployeeAddress
	// emp2Address.Housenumber = 6589
	// emp2Address.Street = "funway st"
	// emp2Address.City = "columbus"
	// emp2Address.State = "ohio"
	// emp2Address.Zipcode = 57932

	// emp2.Address = emp2Address
	employees = append(employees,emp1,emp2)
	fmt.Println(employees)
	StructCheck(employees)
	fmt.Println(employees)
}
func StructCheck([]Employee){
	
	fmt.Println()


}
