package main

import "fmt"

func main() {
	//go execution always starts with func main method ..
	// This is the first golang examples wiht basic control flow
	//It has sequential flow , loop and conditional flow
	fmt.Println("Hello World ... Welcome to  go programming") //sequential
	even()                                                    //sequential
	exit()
}

func even() {
	for count := 0; count <= 100; count++ { //loop
		if count%2 == 0 { //conditional
			fmt.Println(count)
		}
	}
}

func exit() {
	fmt.Println("Leaving bbye~~~~~~~~~~~~")
}
