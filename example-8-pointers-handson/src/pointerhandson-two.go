package main

import "fmt"

func main() {
	first, second, third := 201, 301, 401
	//without pointers
	fmt.Println(first + second + third)
	//assign pointer now
	pt1 := &first  //pt1 is the pointer to first
	pt2 := &second //pt2 is pointer to second
	pt3 := &third  //pt3 is pointer to third
	//sum with pointers
	fmt.Println(*pt1 + *pt2 + *pt3)

	//* denotes pointers underlying value
	//*pt1 is a pointer to pt1 value ,*pt2 is pointer to pt2 value ,*pt3 is pointer to pt3 value
	//fmt.Println(pt1 + pt2 + pt3) - this is invalid no operations supported on pointer itself..Operations are supported on pointer value which is *pt


}
