package maps

import "fmt"

func MapsDefine() {

	mapsofworkers := make(map[string]int)
	mapsofworkers["fieldmanagers"] = 24
	mapsofworkers["team1"] = 12
	mapsofworkers["team2"] = 12
	mapsofworkers["team3"] = 0
	fmt.Println(mapsofworkers)
	fmt.Println(mapsofworkers["team1"])
	fmt.Println(len(mapsofworkers))
	mapsofworkers["fieldmanagers"] = 54
	fmt.Println(mapsofworkers)
	mapsofworkers["Fieldmanagers"] = 674    //maps of case sensitive
	fmt.Println(mapsofworkers)
	fmt.Println(mapsofworkers["team4"])    //unassigned elements will print default value i.e 0
	
	if _, ok := mapsofworkers["team2"]; ok {    //checking if key is present or not (key present "ok")
		fmt.Println("Key is present:",ok)
}
	/*if_, ok := mapsofworkers["team3"]; !ok {  //checking if key is present or not (key not present "!ok")
		fmt.Println("Key is present:", ok)
}*/
	
}