package main

import "fmt"

var first int
var second string
var third bool
func main(){
	fmt.Println("Default Values For Various var Types")
	fmt.Println(first,second,third)
	first=100
	second="Shamsheer Practicing go lang"
	third=true
	s := fmt.Sprint(first," ",second, " " , third)
	fmt.Println(s)
}
