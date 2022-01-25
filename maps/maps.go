package maps

import "fmt"

func MapsDefine() {

	mapsofworkers := make(map[string]int) //map attributes of one entity to attributes of same entity
	mapsofworkers["fieldmanagers"] = 24
	mapsofworkers["team1"] = 12
	mapsofworkers["team2"] = 12
	mapsofworkers["team3"] = 5
	mapsofworkers["team4"] = 0     
	fmt.Println(mapsofworkers)
	fmt.Println(mapsofworkers["team1"])  //accessing one value 
	fmt.Println(len(mapsofworkers))     //lenth of map
	mapsofworkers["fieldmanagers"] = 54  //changing one value in a map
	fmt.Println(mapsofworkers)
	mapsofworkers["Fieldmanagers"] = 674    //maps of case sensitive
	fmt.Println(mapsofworkers)
	fmt.Println(mapsofworkers["team4"])    //unassigned elements will print default value i.e 0
	
  mapsofworkers["team1"] = mapsofworkers["team1"] + 5 //adding 5 to team1
	
	//for range loop is mostly used in maps

	for key,val :=range mapsofworkers { // key and value both are printed
		fmt.Println("key", key)
		fmt.Println("val",val)
		fmt.Println()
	
		mapsofworkers[key]=val+10
	}
	fmt.Println(mapsofworkers)

	for key, _ := range mapsofworkers { // only key printed
		fmt.Println(key)
	}

	for _, val:=range mapsofworkers { // only value printed
		fmt.Println(val)
	}
	fmt.Println(mapsofworkers)
	
	if _, ok := mapsofworkers["team2"]; ok {    //checking if key is present or not (key present "ok")
		fmt.Println("Key is present:",ok)
	
}
	/*if_, ok := mapsofworkers["team3"]; !ok {  //checking if key is present or not (key not present "!ok")
		fmt.Println("Key is present:", ok)
}*/
	val, ispresent := mapsofworkers["team2"]  //value is in val and in ispresent(bool) its checking for key is present or not
	if !ispresent{
		fmt.Println("key is not present")
	}else {
		fmt.Println("key is present")
		fmt.Println(val)
	}
	
     delete(mapsofworkers, "team3")   //deleting a key and value from a map
}