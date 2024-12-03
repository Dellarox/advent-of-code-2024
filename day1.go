package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("inputs/day1.txt")
	check(err)
	dataStr := string(dat)
	dataArr := strings.Fields(dataStr)
	var leftArr []int
	var rightArr []int
	for index, value := range dataArr {
		valueInt, errAtoi := strconv.Atoi(value)
		check(errAtoi)
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
