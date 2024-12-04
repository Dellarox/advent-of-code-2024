package main

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"regexp"
)

func main() {
	dataStr := utilities.ReadFromFileAsOneString("inputs/day3.txt")

	r, _ := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	multiplicationsArr := r.FindAllString(dataStr, -1)

	rCheckNumbers, _ := regexp.Compile(`[0-9]{1,3}`)
	var result int
	for _, mul := range multiplicationsArr {
		numbers := utilities.ConvertStringArrToIntArr(rCheckNumbers.FindAllString(mul, 2))
		result += numbers[0] * numbers[1]
	}
	fmt.Println(result)

	r, _ = regexp.Compile(`don't|do|mul\([0-9]{1,3},[0-9]{1,3}\)`)
	multiplicationsAndInstrArr := r.FindAllString(dataStr, -1)

	isDont := false
	multiplicationsArr = nil
	for _, multOrInstr := range multiplicationsAndInstrArr {
		if multOrInstr == "do" {
			isDont = false
			continue
		}
		if multOrInstr == "don't" {
			isDont = true
			continue
		}
		if !isDont {
			multiplicationsArr = append(multiplicationsArr, multOrInstr)
		}
	}

	var result2 int
	for _, mul := range multiplicationsArr {
		numbers := utilities.ConvertStringArrToIntArr(rCheckNumbers.FindAllString(mul, 2))
		result2 += numbers[0] * numbers[1]
	}
	fmt.Println(result2)
}
