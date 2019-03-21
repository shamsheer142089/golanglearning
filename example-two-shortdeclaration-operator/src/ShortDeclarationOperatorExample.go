package main

import "fmt"

var values=300

func main() {
	//short declaration operator is :=
	shortdeclarationOperatorVariable := 90
	fmt.Println(shortdeclarationOperatorVariable)

	//String Type
	shortdeclarationOperatorSecond := "Shamsheer , Developer"
	fmt.Println(shortdeclarationOperatorSecond)

	shortdeclarationOperatorWithAddition := shortdeclarationOperatorVariable + 300
	fmt.Println(shortdeclarationOperatorWithAddition)
	test()
	fmt.Println(values)
	values=200
	fmt.Println(values)
}

func test(){
	values=500
}