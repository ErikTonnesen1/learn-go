package main

//Arrays

//the type [n]T is an array of n length and of type T
// var a [10]int denotes a as an int array of size 10
//** an array's length is apart of its type ~ so they cannot be resized **

import "fmt"

var names = [4]string{
	"John",
	"Paul",
	"George",
	"Ringo",
}

func arrays_and_slices() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	getSlice(primes)
	fmt.Println("Doing Slice operations: ")
	sliceOperations()
}

//Slices

//a slice is a dynamically sized flexible view into the elements of an array
//slices are more common than arrays

//the type []T is a slice with elements of type T
//a slice is formed by specifying two indices: a low and a high bound, separated by a colon:
// a[low : high]

//**this selects a half-open range which includes the first element but excludes the last one**

//the following expression creates a slice which includes elements 1 thru 3 of a
//a[1:4]

func getSlice(arr [6]int) {
	upperBound := len(arr) - 1
	slice := arr[1:upperBound]
	fmt.Printf("Slice: %v\n", slice)
}

//Slices are like references to arrays
// - a slice does not store any data, it just describes a section of an underlying array
// **- changing the elements of a slice modifies the corresponding elements of its underlying**
// **Other slices that share the same underlying array with see those changes**

func sliceOperations() {
	fmt.Println(names)
	a := names[0:2] //[John, Paul]
	b := names[1:3] //[Paul, George]
	fmt.Println(a, b)

	b[0] = "XXX" //"Paul" = 'XXX'
	fmt.Println(a, b)
	fmt.Println(names)
}

// Slice literals
// like an array literal without the length
