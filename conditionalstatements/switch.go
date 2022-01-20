package conditional

import "fmt"
func SwitchCondition() {
	
	switch a := 10; { // missing expression means "true"
	case a < 12 :
		fmt.Println("a is less than 12")
	case a > 17:
		fmt.Println("a is greater than 17")
	default:
		fmt.Println("default case!")
	}

	switch wow := "wow"; {
	case wow != "wow" :
		fmt.Println("T")
	case wow == "me":
		fmt.Println("B")
	case wow == "wow":
		fmt.Println("C")
	default:
		fmt.Println("default case!")
	}
	switch bool := true; {
    case bool == false :
	fmt.Println("false")
	case bool != true :
	fmt.Println("false")
	default :
	fmt.Println("true")
}
}



	
