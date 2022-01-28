package structures

import "fmt"

type School struct {
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

}