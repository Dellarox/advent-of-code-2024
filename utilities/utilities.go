package utilities

import "strconv"

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
