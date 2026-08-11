/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.

Constraints:

1 <= prices.length <= 105
0 <= prices[i] <= 104
*/
package main

import (
	"fmt"
	"slices"
)

func main() {
	test := []int{7, 1, 5, 3, 6, 4}
	test2 := []int{7, 6, 4, 3, 1}
	test3 := []int{2, 4, 1}
	fmt.Println(maxProfit(test))
	fmt.Println(maxProfit(test2))
	fmt.Println(maxProfit(test3))
}

type profitRange struct {
	max      int
	maxIndex int
	min      int
	minIndex int
}

func maxProfit(prices []int) (maxProfit int) {
	state := profitRange{}
	sorted := slices.Clone(prices)
	slices.Sort(sorted)

	state.min = sorted[0]
	state.minIndex = slices.Index(prices, state.min)

	fmt.Printf("Original Slice: %v\n", prices)
	fmt.Printf("Sorted Slice: %v\n", sorted)
	for i := len(sorted) - 1; i >= 0; i-- {
		sortVal := sorted[i]
		origIndex := slices.Index(prices, sortVal)
		fmt.Printf("Comparing Val: %d, Index: %d\n", sortVal, origIndex)
		if sortVal > state.min && origIndex > state.minIndex {
			fmt.Printf("***Condition Met***: SortVal: %d, Index: %d\n", sortVal, origIndex)
			state.max = sortVal
			state.maxIndex = origIndex
			break
		}
		fmt.Printf("Max: %d, MaxIndex: %d, Min: %d, MinIndex: %d\n", state.max, state.maxIndex, state.min, state.minIndex)
	}

	if maxProfit = state.max - state.min; maxProfit > 0 {
		return
	} else {
		return 0
	}
}
