package main

import "fmt"

func main() {
	var first,second int
	fmt.Println("Please enter two numbers with an enter in between")
	fmt.Scanf("%v\n%v", &first,&second)
	sum:=findSum(first,second)
	fmt.Printf("Sum of two numbers %d and %d  is : %d  ",first,second,sum)
}

func findSum(num1 int,num2 int) int {
	return num1+num2
}