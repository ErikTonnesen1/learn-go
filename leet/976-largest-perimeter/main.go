package main

import (
	"fmt"
	"slices"
)

//Given an integer array nums, return the largest perimeter of a triangle with a non-zero area, formed from three of these lengths. If it is impossible to form any triangle of a non-zero area, return 0.
//
// Example 1:
//
// Input: nums = [2,1,2]
// Output: 5
// Explanation: You can form a triangle with three side lengths: 1, 2, and 2.
// Example 2:
//
// Input: nums = [1,2,1,10]
// Output: 0
// Explanation:
// You cannot use the side lengths 1, 1, and 2 to form a triangle.
// You cannot use the side lengths 1, 1, and 10 to form a triangle.
// You cannot use the side lengths 1, 2, and 10 to form a triangle.
// As we cannot use any three side lengths to form a triangle of non-zero area, we return 0.
//
//
// Constraints:
//
// 3 <= nums.length <= 104
// 1 <= nums[i] <= 106

func main() {

	testSlice1 := []int{1, 2, 3}
	fmt.Println(largestPerimeter(testSlice1))
}

func largestPerimeter(nums []int) (perimeter int) {
	//any two sides of a triangle must add to be greater than the third
	// in other words a + b > c
	// so if we know c, then we need to find the largest a + b combo
	// we can do this by first sorting the array, then going in reverse order, finding the best a + b < c combo

	slices.Sort(nums)

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i]+nums[i-1] > nums[i-2] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return
}
