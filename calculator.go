package calculator

import "fmt"

func CalOperation() {

	x := 20
	y := 30
	z := "play"
	x1 := true

	add := x+y
	sub := x-y


	fmt.Printf("%v,%T \n", add, add)

	fmt.Printf("%v ,%T \n" ,sub, sub)

	fmt.Printf("%v ,%T \n", z, z)

	fmt.Printf("%v ,%T \n", x1, x1)


	
}