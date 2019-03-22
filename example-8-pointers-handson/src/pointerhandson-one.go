package main

import "fmt"

func main() {
	var pt *int
	value := 500
	//when we assign variable address to pointer , *pt returns the value at address
	//value is 500 and its address assigned to a pointer 'pt' ..when we print *pt it will print 500
	pt = &value
	fmt.Println(*pt)
	fmt.Println(&pt)
	fmt.Println(&value)
	//when we assign pointers address to another pointer , pointer stores the exact address of a pointer value
	//i.e, &value == *pt4 in below example
	//the below two lines demonstrate this statement
	//pt is pointer to value ,
	pt4 := &pt
	//*pt4=400 - if we try to assign the some integer value to *pt4 which actually contains the address of another pointer
	//It straight away throws compile time error.
	//If we really want to store a value to *pt4 , we need refer the pointer of the pointer , that is we need to add one more * before pointer
	//eg **pt4=400
	//**pt4=400
	fmt.Println(*pt4)  //it prints the address of pt
	fmt.Println(**pt4) //it prints the value at pt

	//Let us store address of pt4 in another pointer
	pt5:=&pt4
	fmt.Println("Address stored at pt4(which is of pt in this example) is the value in pt5 -->",pt4,*pt5)
	fmt.Println("Address stored at pt(which is of value variable in this example) is the value in **pt5 -->",pt,**pt5)
	fmt.Println("Value stored at value variable is the value in ***pt5 --> ",value,***pt5)




}
