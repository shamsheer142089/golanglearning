package main

import "fmt"

func main() {
	//declaring map - method one
	movies := map[string]string{
		"Mic":   "100",
		"Mouse": "200",
	}

	fmt.Println(movies)

	//declaring map - method two -Zero value for map is just empty map

	var cost map[string]string
	cost["id1"]="300"
	fmt.Println(cost)

	//declaring map - method three -Zero value for map is just empty map

	var price=make(map[string]string)
	price["id2"]="400"
	fmt.Println(price)
	delete(price,"id2")
}
