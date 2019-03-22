package main

import "fmt"

func main() {
	var price=make(map[string]string)

	//declaring map - method three -Zero value for map is just empty map

	price["id2"]="400"
	price["id3"]="500"
	price["id6"]="600"
	fmt.Println(price)
	iterateAndPrintMap(price)
	delete(price,"id2")
}

func iterateAndPrintMap(myMap map[string]string){
	for key,value:= range myMap{
		fmt.Println("The value of key --> ",key," Is --> ",value)
	}
}
