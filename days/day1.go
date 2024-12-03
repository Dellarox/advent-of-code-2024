package main

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("inputs/day1.txt")
	utilities.Check(err)
	dataStr := string(dat)
	dataArr := strings.Fields(dataStr)
	var leftArr []int
	var rightArr []int
	for index, value := range dataArr {
		valueInt, errAtoi := strconv.Atoi(value)
		utilities.Check(errAtoi)
		if index%2 == 0 {
			leftArr = append(leftArr, valueInt)
		} else {
			rightArr = append(rightArr, valueInt)
		}
	}

	sort.Ints(leftArr)
	sort.Ints(rightArr)

	// First solution
	var finalResult int
	for index, value := range leftArr {
		finalResult += int(math.Abs(float64(value - rightArr[index])))
	}
	fmt.Println(finalResult)

	numberOfNumbers := make(map[int]int)

	for _, valueUniqueLeft := range leftArr {
		for _, valueRight := range rightArr {
			if valueUniqueLeft == valueRight {
				numberOfNumbers[valueUniqueLeft]++
			}
		}
	}
	fmt.Println(numberOfNumbers)

	// Second solution
	var finalResult2 int
	for number, numberCount := range numberOfNumbers {
		finalResult2 += numberCount * number
	}
	fmt.Println(finalResult2)
}
