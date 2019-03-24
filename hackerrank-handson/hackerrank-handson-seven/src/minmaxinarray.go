package main

import "fmt"

func main() {
	var arr []int64
	arr = append(arr, 769082435,210437958,673982045,375809214,380564127)
	miniMaxSum(arr)
}

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int64) {
	var sumArr []int64
	sumOfTotalElementsInArray := findSum(arr)
	for _, value := range arr {
		var totalExceptCurrent int64
		totalExceptCurrent = sumOfTotalElementsInArray - value
		sumArr = append(sumArr, totalExceptCurrent)
	}

	min, max := findMinimumAndMaximumFromArray(sumArr)
	fmt.Println(min, max)
}

func findMinimumAndMaximumFromArray(array []int64) (int64, int64) {
	var minVal int64
	var maxVal int64
	minVal = array[0]
	maxVal = array[0]
	for _, minOrMax := range array {
		if minOrMax > maxVal {
			maxVal = minOrMax
		}

		if minVal > minOrMax {
			minVal = minOrMax
		}
	}
	return minVal, maxVal
}

func findSum(arri []int64) int64 {
	var sum int64 = 0
	for _, value := range arri {
		sum = sum + value
	}
	return sum
}
