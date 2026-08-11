package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	k := 23
	ans := smallestRepunitDivByK(k)
	fmt.Printf("Smallest divisible number of %d has a length of %d\n", k, ans)
}

// N may NOT fit in a int64 int
// build # by adding 1's and seeing if divisible
// how to tell if not divisible by K?
func smallestRepunitDivByK(k int) int {
	var num int = 1
	for num > 0 {
		fmt.Printf("Iteration: %d\n", num)
		if num%k == 0 {
			strNum := strconv.Itoa(num)
			fmt.Println(num)
			return len(strNum)
		}
		num = (num*10 + 1)
	}
	return -1
}
