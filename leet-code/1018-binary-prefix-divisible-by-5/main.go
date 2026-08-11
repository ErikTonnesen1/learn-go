package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	nums := []int{1, 0, 1, 1, 0}
	answers := prePrefixesDivBy5(nums)
	fmt.Println("Answer to answers = ", answers)
	answers2 := prefixesDivBy5(nums)
	fmt.Println("Answer to answers = ", answers2)

}

// My solution before seeing that bitwise manipulation is the correct answer
func prePrefixesDivBy5(nums []int) (answers []bool) {
	answers = make([]bool, len(nums))
	for i := range nums {
		binaryArray := nums[0 : i+1]
		binaryString := joinSliceItoa(binaryArray)
		binaryInt, err := strconv.ParseInt(binaryString, 2, 64)
		// fmt.Printf("Binary String: %s, Binary Representation: %d\n", binaryString, binaryInt)
		if err != nil {
			log.Fatal("Binary Decode broke", err)
		}

		if binaryInt%5 == 0 {
			answers[i] = true
		} else {
			answers[i] = false
		}
	}
	return answers
}

func joinSliceItoa(nums []int) (sliceAsString string) {
	for _, v := range nums {
		sliceAsString += strconv.Itoa(v)
	}
	return sliceAsString
}

func prefixesDivBy5(nums []int) (answers []bool) {
	answers = make([]bool, len(nums))
	var num int
	for i, v := range nums {
		//shifts to the left by 1 (appends 0 onto binary #)
		// checks if v
		num = ((num<<1 | v) % 5)
		answers[i] = num == 0
	}
	return answers
}
