package main

import (
	"advent-of-code-2024/utilities"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

func convertStringArrToIntArr(strArr []string) []int {
	var intArr []int
	for _, elem := range strArr {
		intElem, err := strconv.Atoi(elem)
		utilities.Check(err)
		intArr = append(intArr, intElem)
	}
	return intArr
}

func main() {
	file, err := os.Open("inputs/day2.txt")
	utilities.Check(err)
	scanner := bufio.NewScanner(file)
	var fullDataArr [][]string
	for scanner.Scan() {
		lineArr := strings.Fields(scanner.Text())
		fullDataArr = append(fullDataArr, lineArr)
	}
	var countTrues1 int
	for _, report := range fullDataArr {
		intReport := convertStringArrToIntArr(report)
		safeStatement := isSafe(intReport)
		if safeStatement {
			countTrues1++
		}
	}

	var countTrues2 int
	for _, report := range fullDataArr {
		for i := range len(report) {
			intReport := convertStringArrToIntArr(report)
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
