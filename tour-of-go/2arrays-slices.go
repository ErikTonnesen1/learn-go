package main

//Arrays

//the type [n]T is an array of n length and of type T
// var a [10]int denotes a as an int array of size 10
//** an array's length is apart of its type ~ so they cannot be resized **

import (
	"fmt"
	"strings"
)

var names = [4]string{
	"John",
	"Paul",
	"George",
	"Ringo",
}

func slices_and_arrays() {
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

	fmt.Println("\nDoing Slice literal stuff")
	sliceLiteral()

	fmt.Println("\nDoing Slice literal stuff")
	sliceDefaults()

	fmt.Println("\nDoing Slice length and capacity stuff:")
	sliceLengthAndCapacity()

	fmt.Println("\nDoing Nil Slice stuff")
	nilSlice()

	fmt.Println("\nMaking Slices with the make method")
	makeSlices()

	fmt.Println("\nDoing slices containing slices stuff")
	sliceOfSlices()

	fmt.Println("\nDoing Appending slices stuff")
	appendSlice()
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

//Array Literal
//Example: [3]bool{true, true, false}

//Slice literal
//this creates the same array as above then builds a slice that references it
//[]bool{true, true, false}

func sliceLiteral() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{ //slice literal is of type struct that is defined inline
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(s)
}

// Slice defaults
//instead of including the lower and upper bounds in a slice definition, you can omit these and it will
//use their default values:
//	- lower bound = 0
//	- upper boud = length of slice

//for the array
// var a [10]int
// these slices are equivalent:

// a[0:10]
// a[:10]
// a[0:]
// a[:]

func sliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4] //[3, 5, 7]
	fmt.Println(s)

	s = s[:2] //slice of the slice [3, 5]
	fmt.Println(s)

	s = s[1:] //slice of the sliced slice: [5]
	fmt.Println(s)

}

//Slice length and capacity
// A slice has both a length and a capacity
// - length: # of elements it contains
// - capacity: number of elems in the underlying array,
//		*counting from first element in the slice*

// length and capacity of the slice can be obtained from the following methods
// - len(s) -> length
// - cap(s) -> capacity

//**You can extend a slice's length by re-slicing it, given it has capacity
//In the following method, we'll try and extend the length beyond capacity
//and see what happens

func sliceLengthAndCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	//slice the slice to give it length 0
	s = s[:0]
	printSlice(s)

	//extending length
	s = s[:4]
	printSlice(s)

	//Drop first two values
	s = s[2:]
	printSlice(s)
}

func printSlice(slice []int) {
	fmt.Printf("len=%d, cap=%d, %v\n", len(slice), cap(slice), slice)
}

//nil slices
// the zero value of a slice is *nil*
// a nil slice has length and capacity = 0 and no underlying array

func nilSlice() {
	var s []int
	printSlice(s)
	if s == nil {
		fmt.Println("nil!")
	}
}

// Creating Slices with 'make' a.k.a ~dynamically~ sized arrays

// - 'make' allocates a zeroed array and returns a slice that refers to that
//	array

// a := make([]int, 5) // len(a) = 5

// to specify a capacity, pass a 3rd arg

// b := make([]int, 0, 5) // len(b) = 0, cap(b) = 5

// b = b[:cap(b)] // len(b)=5, cap(b)=5
// b = b[1:] //len(b)=4, cap(b)=4,

func makeSlices() {
	a := make([]int, 5)
	printMakeSlices("a", a) // len=5, cap=5

	b := make([]int, 0, 5)
	printMakeSlices("b", b) // len=0, cap=5

	c := b[:2]
	printMakeSlices("c", c) //len=2 cap=5

	d := c[2:5]
	printMakeSlices("d", d) //len=3, cap=3
}

func printMakeSlices(s string, x []int) {
	fmt.Printf("%s, len=%d, cap=%d %v\n", s, len(x), cap(x), x)
}

//Slices can contain any type, including other slices

func sliceOfSlices() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][0] = "O"
	board[0][1] = "X"
	board[2][2] = "O"
	board[1][0] = "X"
	board[2][2] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Println(strings.Join(board[i], " "))
	}
}

// Appending slices
//you can append items to a slice via the following method from Go's 'builtin' standard libary
//func append(slice []T, ... T)
// - first param is a slice
// - second param(s) are the items to append to the slice

// append checks to see if the underlying array of the slice has capacity to add to it, if not, a new underlying array is allocated, and a new pointer to the new array is returned by append

func appendSlice() {
	var s []int //nil slice
	printSlice(s)

	//append works on nil slices
	s = append(s, 0)
	printSlice(s)

	//slice continues to grow
	s = append(s, 1)
	printSlice(s)

	// can append mulitple items at once
	s = append(s, 2, 3, 4, 5)
	printSlice(s)
}
