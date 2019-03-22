package main

import "fmt"

func main() {
	numberSlice:=[]int{0,1,2,3,4,5,6,7,9,10}

	for _,number:= range  numberSlice{
		if number % 2 == 0 {
			fmt.Printf("%d is even \n" , number)
		}else{
			fmt.Printf("%d is odd \n" , number)
		}

	}

}
