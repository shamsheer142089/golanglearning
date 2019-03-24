package main

import (
	"fmt"
	"math"
)

// Complete the diagonalDifference function below.
func diagonalDifference(arr [4][4]int32) int32 {
	var leftDiagonalSum int32 = 0
	var rightDiagonalSum int32 = 0
	fmt.Println("length", len(arr))
	for idx, _ := range arr {
		for idx2, _ := range arr {
			if idx == idx2 {
				leftDiagonalSum = leftDiagonalSum + arr[idx][idx2]
			}
			if idx+idx2 == len(arr)-1 {
				fmt.Println("right", arr[idx][idx2])
				rightDiagonalSum = rightDiagonalSum + arr[idx][idx2]
			}

		}
	}
	fmt.Println(leftDiagonalSum)
	fmt.Println(rightDiagonalSum)
	return absoluteValue(leftDiagonalSum - rightDiagonalSum)
}

func absoluteValue(n int32) int32 {
	return int32(math.Abs(float64(n)))
}

func main() {
	var arr [4][4]int32
	var k int32 = 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			arr[i][j] = k
			k = k + 1
		}
	}
	fmt.Println(arr)
	result := diagonalDifference(arr)

	fmt.Println(result)
}
