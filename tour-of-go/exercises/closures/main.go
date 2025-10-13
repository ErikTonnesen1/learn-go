package main

//implement a fibonacci function that returns a function(a closure) that returns successive fibonacci numbers
// 0, 1, 1, 2, 3, 5, 8, 11, ...

import "fmt"

func fibonacci() func() int {
	n2 := 0
	n1 := 1
	iteration := 0
	return func() int {
		if iteration < 2 {
			setup := iteration
			iteration++
			return setup

		}
		n := n2 + n1
		n2 = n1
		n1 = n
		return n
	}
}

func main() {
	f := fibonacci()
	for range 10 {
		fmt.Println(f())
	}

}
