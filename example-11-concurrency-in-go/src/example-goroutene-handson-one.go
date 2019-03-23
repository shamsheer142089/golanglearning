package main

import "fmt"

func main() {
	for st:=0;st<=10;st++{
		go evenOrOdd(st)  // go key word before the function makes it as routine
	}
}

func evenOrOdd(val int){
	if val%2==0{
		fmt.Println("Iam Even : ",val)
	}else{
		fmt.Println("Iam Odd : ",val)
	}
}

