package main

import "fmt"

var number int
var content string
var decimals32 float32
var decimal64 float64
var boolval bool

func main() {

	fmt.Println(number)
	fmt.Printf("%T \n",number)
	fmt.Println(content)
	fmt.Printf("%T \n",content)
	fmt.Println(decimals32)
	fmt.Printf("%T \n",decimals32)
	fmt.Println(decimal64)
	fmt.Printf("%T \n",decimal64)
	fmt.Println(boolval)
	fmt.Printf("%T \n",boolval)
}
