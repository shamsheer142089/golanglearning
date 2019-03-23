package main

import "fmt"

func main() {
	first, second := 201, 301
	fmt.Println("Numbers Before Swapping -->", first, second)
	swappedFirstNumber, swappedSecondNumber := swapGeneralway(first, second)
	fmt.Println("Numbers After Swapping -->", swappedFirstNumber, swappedSecondNumber)
	swappedOne,swappedTwo:=swapGeneralwayWithPointers(&first, &second)
	fmt.Println("Numbers Swapped Using Pointers --> ",swappedOne," ",swappedTwo)
	//another easy way to swap in GO . go by default uses registry to store temporary values.
	//so we dont need to use temporary variable
	fmt.Println("Numbers before swapping in go way",first,second)
	second,first=first,second
	fmt.Println("Numbers after swapping in go way",first,second)
}

func swapGeneralway(first int, second int) (int, int) {
	tempValue := first
	first = second
	second = tempValue
	return first, second
}

func swapGeneralwayWithPointers(numb1 *int, numb2 *int) (int, int) {
	tempValue := *numb1
	*numb1 = *numb2
	*numb2 = tempValue
	return *numb1, *numb2
}
