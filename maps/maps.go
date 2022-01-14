package maps

import "fmt"

func MapsDefine() {

	mapsofworkers := make(map[string]int)
	mapsofworkers["fieldmanagers"] = 12
	mapsofworkers["team1"] = 12
	mapsofworkers["team2"] = 12

	fmt.Println(mapsofworkers)


}