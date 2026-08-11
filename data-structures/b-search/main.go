package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 5, 6, 8, 9}
	// binarySearch(arr, 0, (len(arr) - 1), 11)
	fmt.Println(bSearch(arr, 4))
}

// first go around
func binarySearch(arr []int, lb, ub, x int) {
	if ub < lb || lb >= len(arr) || ub >= len(arr) {
		fmt.Println("-1")
	}

	median := lb + (ub-lb)/2

	if arr[median] == x {
		fmt.Printf("%d\n", median)
	}

	if arr[median] > x {
		ub = median - 1
		binarySearch(arr, lb, ub, x)
	}
	if arr[median] < x {
		lb = median + 1
		binarySearch(arr, lb, ub, x)
	}
}

// correct go-form
func bSearch(arr []int, search int) (result, searchCount int) {
	mid := len(arr) / 2
	switch {
	case len(arr) == 0:
		result = -1
	case arr[mid] > search:
		result, searchCount = bSearch(arr[:mid], search)
	case arr[mid] < search:
		result, searchCount = bSearch(arr[mid:], search)
	default:
		result = mid
	}
	searchCount++
	return
}
