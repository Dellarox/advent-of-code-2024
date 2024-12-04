package utilities

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func ConvertStringArrToIntArr(strArr []string) []int {
	var intArr []int
	for _, elem := range strArr {
		intElem, err := strconv.Atoi(elem)
		Check(err)
		intArr = append(intArr, intElem)
	}
	return intArr
}

func ReadFromFileAsOneString(filePath string) string {
	dat, err := ioutil.ReadFile(filePath)
	Check(err)
	return string(dat)
}

func ReadFromFileLineByLine(filePath string) [][]string {
	file, err := os.Open(filePath)
	Check(err)
	scanner := bufio.NewScanner(file)
	var fullDataArr [][]string
	for scanner.Scan() {
		lineArr := strings.Fields(scanner.Text())
		fullDataArr = append(fullDataArr, lineArr)
	}
	return fullDataArr
}
