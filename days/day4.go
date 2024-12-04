package main

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"strings"
)

func main() {
	fullDataArr := utilities.ReadFromFileLineByLine("inputs/day4.txt")
	var correctDataArr [][]string
	fmt.Println(fullDataArr[1])
	for _, line := range fullDataArr {
		lineArr := strings.Split(line[0], "")
		correctDataArr = append(correctDataArr, lineArr)
	}
	fmt.Println(correctDataArr[1][6])
}
