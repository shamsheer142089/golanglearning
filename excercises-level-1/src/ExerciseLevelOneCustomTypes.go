package main

import "fmt"

type burgerPrice int
type burgerType string
var customValue burgerPrice
var customType burgerType
var burgerCost int
func main(){
	//custom type examples
	fmt.Println(customValue)
	fmt.Printf("%T \n",customValue)
	fmt.Printf("%T \n",customType)
	customValue=42
	customType="Zinger Chicken"
	fmt.Println(customValue)
	fmt.Printf("%T \n",customValue)
	fmt.Println(customType)
	fmt.Printf("%T \n",customType)
	burgerCost=30
	burgerCost=int(customValue)
	newValue:=int(customValue)
	fmt.Println(newValue)
	fmt.Printf("%T \n",burgerCost)

	//type conversion
	


}
