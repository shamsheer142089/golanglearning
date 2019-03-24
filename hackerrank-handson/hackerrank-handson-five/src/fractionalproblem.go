package main

import (
	"fmt"
	)

// Complete the diagonalDifference function below.
// Complete the plusMinus function below.
func plusMinus(arr []int32) {
	var posCount int32 = 0
	var negCount int32 = 0
	var zeroCount int32 = 0
	totalCount := len(arr)
	for _, val := range arr {
		if val > 0 {
			posCount = posCount + 1

		} else if val < 0 {
			negCount = negCount + 1
		} else if val == 0 {
			zeroCount = zeroCount + 1
		}

	}
	fmt.Printf("%.6f \n", getRatio(posCount, int32(totalCount)))
	fmt.Printf("%.6f \n", getRatio(negCount, int32(totalCount)))
	fmt.Printf("%.6f \n", getRatio(zeroCount, int32(totalCount)))

}

func getRatio(frac int32, total int32) float64 {
	return float64(frac) / float64(total)
}

func main() {
	var arr []int32
	arr = append(arr, -1)
	arr = append(arr, 1)
	arr = append(arr, -1)
	arr = append(arr, 31)
	arr = append(arr, 0)
	arr = append(arr, -1)
	arr = append(arr, -7)
	fmt.Println(arr)

	plusMinus(arr)

}
