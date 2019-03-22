package main

import "fmt"

func main() {
	movieName:=getMovieName()
	fmt.Println(movieName)
	printMovieName() //this function is defined in functionreturntypehelper class run these
}

func getMovieName() string{
	return "Tom And Jerry Tales"
}