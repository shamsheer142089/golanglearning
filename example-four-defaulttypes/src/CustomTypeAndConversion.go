package main

import "fmt"

var number int
//chocolate is custom type and string is its actual type
type chocolate string
type chocolatecost int
func main() {
	chocolate := "Five Star"
	number=300
	chocolatecost := 500
	fmt.Println(chocolatecost)
	// chocolatecost=number  - this is not possible
	// in golan because chocolatecost is a different type and number is diffrent type.
	// This is called casting and its not possible in go land instead we can convert types
	chocolatecost =int(number)
	fmt.Println(chocolate)

}
