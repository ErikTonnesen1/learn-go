package main

import (
	"fmt"
	"slices"
)

func main() {
	arr := []int{2, 2, 3, 4}
	arr2 := []int{1, 2, 2, 3, 3, 3}
	arr3 := []int{2, 2, 2, 3, 3}

	fmt.Println(findLucky(arr))
	fmt.Println(findLucky(arr2))
	fmt.Println(findLucky(arr3))
}

//go maps are very interesting
// maps have the syntax map[keyType]valueType
// to create a map, use the make function := make(map[int]string)
// to GET a value-> i, ok := map[1]
// 	- i is the returned value of the get call, if there exists 1 in the map, the value of that k,v pair is returned. if not, then the default zero value is returned (int: 0, string: "", bool: false)
// 	- ok is a bool that is also the answer if the value was found in the map (if found = true, else false)
// 	- can just use the bool if not wanting the value : _, ok := map[1]
// to delete a value in the map, use the delete function ( does not return anything ) delete(m, "route") (deletes "route" in the map 'm' will do nothing if "route" does not exist)

// runtime of 4ms -> so bad
func findLucky1(arr []int) int {
	slices.Sort(arr)
	numMap := make(map[int]int)
	for _, v := range arr {
		_, ok := numMap[v]
		if ok {
			numMap[v] = numMap[v] + 1
		} else {
			numMap[v] = 1
		}

	}
	luckyNumberArray := make([]int, len(numMap))
	for _, v := range numMap {
		if numMap[v] == v {
			luckyNumberArray = append(luckyNumberArray, v)
		}
	}

	largestLuckyNum := slices.Max(luckyNumberArray)

	if largestLuckyNum > 0 {
		return largestLuckyNum
	} else {
		return -1
	}
}

//time complexity = n2

//better attempt
//- runtime of 0ms
//- time complexity of O(n)
//- space complexity of O(n) - stateArray

// use a slice as a way to maintain state
// check if the value in arr is < the length of the array, as if it is longer, it surely could not be a lucky num
func findLucky(arr []int) int {
	size := len(arr)
	stateArray := make([]int, size+1)
	for _, v := range arr {
		if v <= size {
			stateArray[v]++
		}
	}
	//at this point all items in array have been counted, time to find the largest
	for i := size; i >= 1; i-- {
		if stateArray[i] == i {
			return i
		}
	}
	return -1
}
