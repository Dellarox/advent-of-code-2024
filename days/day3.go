package main

import (
	"advent-of-code-2024/utilities"
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	dat, err := ioutil.ReadFile("inputs/day3.txt")
	utilities.Check(err)
	dataStr := string(dat)

	r, _ := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	multipliactionsArr := r.FindAllString(dataStr, -1)

	r, _ = regexp.Compile(`[0-9]{1,3}`)
	var result int
	for _, mul := range multipliactionsArr {
		numbers := utilities.ConvertStringArrToIntArr(r.FindAllString(mul, 2))
		result += numbers[0] * numbers[1]
	}
	fmt.Println(result)
}
