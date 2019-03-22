package main

import "fmt"

func main() {
	movieTitles:=[] string {"Titanic","Tom And Jerry Tales",getMovieName()} //here movieTitles is slice
	movieTitles=append(movieTitles,"Jurassic Park") //added new element to slice
	//Iteration
	for index,movieTitle:=range movieTitles{
		fmt.Println(index,"-----",movieTitle)
	}

}

func getMovieName() string{
	return "Blue Sea"
}