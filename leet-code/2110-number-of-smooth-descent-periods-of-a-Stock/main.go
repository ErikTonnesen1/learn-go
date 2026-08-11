package main

import (
	"fmt"
)

//find all desc ordered patterns length > 1
//add length of array at end to account for single day descents

func main() {
	testCase := []int{3, 2, 1, 4}
	fmt.Println(getDescentPeriods(testCase))

	testCase2 := []int{8, 6, 7, 7}
	fmt.Println(getDescentPeriods(testCase2))

	testCase3 := []int{1}
	fmt.Println(getDescentPeriods(testCase3))
}

// first day is exempt
func getDescentPeriods(prices []int) int64 {
	var numberOfDescents int = 1
	var descPeriod int = 1

	for i := 1; i < len(prices)-1; i++ {
		if prices[i] == prices[i-1]-1 {
			descPeriod++
		} else {
			descPeriod = 1
		}
		numberOfDescents += descPeriod
	}
	return int64(numberOfDescents)
}
