package main

import "fmt"

func main() {
	var arr []int64
	arr = append(arr, 769082435, 210437958, 769082435, 375809214, 380564127)
	fmt.Println(birthdayCakeCandles(arr))
}

// Complete the birthdayCakeCandles function below.
func birthdayCakeCandles(ar []int64) int64 {
	var candleCount int64 = 0
	max := findMaximumFromArray(ar)
	for _, value := range ar {
		if max-value == 0 {
			candleCount = candleCount + 1
		}
	}
	return candleCount
}

func findMaximumFromArray(array []int64) int64 {
	var maxVal int64
	maxVal = array[0]
	for _, minOrMax := range array {
		if minOrMax > maxVal {
			maxVal = minOrMax
		}
	}
	return maxVal
}
