package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(timeConversion("12:00:00 PM"))
}

func timeConversion(s string) string {
	/*
	 * Write your code here.
	 */
	sample12hrsTimeformat := "3:04:05 PM"
	sample24hrsTimeformat := "15:04:05 PM"
	convertedTime, error := time.Parse(sample12hrsTimeformat, s)

	if error != nil {
		fmt.Println("Error while parsing the time literals", error)

	}
	return convertedTime.Format(sample24hrsTimeformat)

}
