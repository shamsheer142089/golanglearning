# golanglearning
golangpractices

Go Pointers
Go language has pointers. A pointer holds the memory address of the value.

var pt *int -> pt is a pointer

value:=200

pt=&value , pt stores the address of value , *pt returns value 200

The & operator generates a pointer to its operand.
The * operator denotes the pointer's underlying value

Few Important Notes With Code Snippers

`
    
    
    var pt *int
    
	value := 500
	_//when we assign variable address to pointer , *pt returns the value at address
	//value is 500 and its address assigned to a pointer 'pt' ..when we print *pt it will print 500_
	pt = &value
	fmt.Println(*pt)
	fmt.Println(&pt)
	fmt.Println(&value)
	_//when we assign pointers address to another pointer , pointer stores the exact address of a pointer value
	//i.e, &value == *pt4 in below example
	//the below two lines demonstrate this statement
	//pt is pointer to value ,_
	pt4 := &pt
	_//*pt4=400 - if we try to assign the some integer value to *pt4 which actually contains the address of another pointer
	//It straight away throws compile time error.
	//If we really want to store a value to *pt4 , we need refer the pointer of the pointer , that is we need to add one more * before pointer
	//eg **pt4=400
	//**pt4=400_
	fmt.Println(*pt4) //it prints the address of pt
	fmt.Println(**pt4) //it prints the value at pt
`





