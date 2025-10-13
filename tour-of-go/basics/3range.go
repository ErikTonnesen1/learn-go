package main

//you can iterate over a slice or array using a for loop in the ~range~ form.
//when "ranging" over a slice, two values are returned for each iteration,
// - the first is the index
// - the second is a ~copy~ of the element at that index

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func rangeFunc() {
	Range(pow)

	fmt.Println("\nOmit index or value")
	omitRangeVars()
}

func Range(slice []int) {
	for i, v := range slice {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

//You can skip the index or value by assigning to _
// for i, _ := range pow
//for _, value := range pow

//if you only want the index, you can omit the second variable
//for i := range pow

func omitRangeVars() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i // this is so cool btw
	}

	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}
