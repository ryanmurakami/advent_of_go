package main

import (
	"fmt"
)

func main() {
	start := 265275
	stop := 781584

	validPasses := 0

	for i := start; i <= stop; i++ {
		if doesMeet(i) {
			validPasses++
		}
	}

	fmt.Println(validPasses)
}

func doesMeet(pass int) bool {
	passArray := intToArray(pass)
	duplicates := make([]int, 9)
	prevDigit := 0
	for _, digit := range passArray {
		if digit < prevDigit {
			return false
		} else if digit == prevDigit {
			duplicates[digit-1]++
		}
		prevDigit = digit
	}
	fmt.Printf("pass %v had %v duplicates\n", pass, duplicates)

	//has 2 digit match
	for _, duplicate := range duplicates {
		if duplicate == 1 {
			return true
		}
	}
	return false
}

func intToArray(pass int) []int {
	intPass := make([]int, 6)
	intPass[5] = pass % 10
	intPass[4] = pass % 100 / 10
	intPass[3] = pass % 1000 / 100
	intPass[2] = pass % 10000 / 1000
	intPass[1] = pass % 100000 / 10000
	intPass[0] = pass % 1000000 / 100000
	return intPass
}
