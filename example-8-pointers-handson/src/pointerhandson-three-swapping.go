package main

import "fmt"

func main() {
	first, second := 201, 301
	fmt.Println("Numbers Before Swapping -->", first, second)
	swappedFirstNumber, swappedSecondNumber := swap(first, second)
	fmt.Println("Numbers After Swapping -->", swappedFirstNumber, swappedSecondNumber)
	swappedOne,swappedTwo:=swapWithPointers(&first, &second)
	fmt.Println("Numbers Swapped Using Pointers --> ",swappedOne," ",swappedTwo)
}

func swap(first int, second int) (int, int) {
	tempValue := first
	first = second
	second = tempValue
	return first, second
}

func swapWithPointers(numb1 *int, numb2 *int) (int, int) {
	tempValue := *numb1
	*numb1 = *numb2
	*numb2 = tempValue
	return *numb1, *numb2
}
