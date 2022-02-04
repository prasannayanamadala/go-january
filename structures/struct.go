package structures

import "fmt"

/*type School struct {          //structures are man defined datatypes
	Name      		string
	yearofjoining   int
	students   		[]string
}

func StructPractice(){
	var w1 School
	w1.Name = "Monon trail"
	
	w1.students = make([]string,0)
	w1.students = append(w1.students,"Amelia","joe","valentina")

	w1.yearofjoining = 2022

	fmt.Println("School:",w1)
	fmt.Println("School Name:",w1.Name)
	fmt.Println("Name of students:",w1.students)
	fmt.Println("year of joining:",w1.yearofjoining)
	fmt.Println("name of first student:",w1.students[0])
	fmt.Println()

	var w2 School
	w2.Name = "west washington"
	
	w2.students = make([]string,0)
	w2.students = append(w2.students,"Amelia","joe","valentina")

	w2.yearofjoining = 2022

	fmt.Println("School:",w2)
	fmt.Println("School Name:",w2.Name)
	fmt.Println("Name of students:",w2.students)
	fmt.Println("year of joining:",w2.yearofjoining)
	fmt.Println("name of first student:",w2.students[0])
	w2.students = append(w2.students,"jackson")
	fmt.Println("Name of students:",w2.students)
}*/
type Employee struct {
	Name string
	Salary int
	Position string
	Address EmployeeAddress
}
type EmployeeAddress struct {
	Housenumber int
	Street string
	City string
	State string
	Zipcode int
}
 func EmployeeDefine() {
	 employees := make([]Employee,0)
	
	var emp1 Employee
	emp1.Name = "jenny"
	emp1.Salary = 25000
	emp1.Position = "Designer"
 
	var emp1Address EmployeeAddress
	emp1Address.Housenumber = 638
	emp1Address.Street = "woodbury ct"
	emp1Address.City = "saltlake city"
	emp1Address.State = "utah"
	emp1Address.Zipcode = 56768

	emp1.Address = emp1Address
	
	//employees = append(employees,emp1)
	//fmt.Println(employees)

	var emp2 Employee
	emp2.Name = "mary"
	emp2.Salary = 30000
	emp2.Position = "cashier"

	var emp2Address EmployeeAddress
	emp2Address.Housenumber = 6589
	emp2Address.Street = "funway st"
	emp2Address.City = "columbus"
	emp2Address.State = "ohio"
	emp2Address.Zipcode = 57932

	emp2.Address = emp2Address
	employees = append(employees,emp1,emp2)
	fmt.Println(employees)

	//for range loop
	
	for i,_ := range employees{ //for index
		//fmt.Println(val)
		fmt.Println(i)
	}

	for _,val := range employees{ //for val
		fmt.Println(val)
		//fmt.Println(i)
	}
	
}

