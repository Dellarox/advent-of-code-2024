package main

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"math"
)

func isSafe(report []int) bool {
	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(report); i++ {
		difference := math.Abs(float64(report[i] - report[i-1]))
		if difference > 3 || difference < 1 {
			return false
		}
		if report[i] > report[i-1] {
			isDecreasing = false
		}
		if report[i] < report[i-1] {
			isIncreasing = false
		}
	}
	return isIncreasing || isDecreasing
}

func removeElemFromSlice(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	fullDataArr := utilities.ReadFromFileLineByLine("inputs/day2.txt")
	var countTrues1 int
	for _, report := range fullDataArr {
		intReport := utilities.ConvertStringArrToIntArr(report)
		safeStatement := isSafe(intReport)
		if safeStatement {
			countTrues1++
		}
	}

	var countTrues2 int
	for _, report := range fullDataArr {
		for i := range len(report) {
			intReport := utilities.ConvertStringArrToIntArr(report)
			intReport = removeElemFromSlice(intReport, i)
			safeStatement := isSafe(intReport)
			if safeStatement {
				countTrues2++
				break
			}
		}
	}

	fmt.Println(countTrues1)
	fmt.Println(countTrues2)
}
