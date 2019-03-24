package main

import "fmt"

func main() {
	staircase(10)
}

// Complete the staircase function below.
func staircase(n int32) {
	var idx int32
	var jdx int32
	var p int32 = 1
	var ps int32
	var pdx int32
	ps = n - 1
	for idx = 0; idx < n; idx++ {
		for jdx = 0; jdx < ps; jdx++ {
			fmt.Printf(" ")
		}

		for pdx = 0; pdx < p; pdx++ {
			fmt.Printf("#")
		}

		ps--
		p++
		fmt.Printf("\n")
	}

}
