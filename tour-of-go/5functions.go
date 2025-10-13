package main

//functions are values too, they can be passed around just like other values
//function values may be used as function arguments and return values

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

var hypot = func(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func main() {
	fmt.Println("HEllo world")
	fmt.Println(compute(hypot)) // can pass functions as args to "compute" which will always run the function with the args 3, 4
	fmt.Println(compute(math.Pow))
}

//Go Functions may be closures
//Closures are function values that access variables outside its body
// - the function may access and assign to the referenced variables
// - in this sense, the function is 'bound' to the variables

//i.e. the 'adder' function returns a closure, Each closure is bound to its own 'sum' variable

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func closure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

//Closures are SUPER interesting, in essence you are creating and returning a function with STATE
// each time a closure method is called, and it returns a function that references a var outside it's body, that var is only accessible
// thru the returned function. e.g. if sum is only accessed through the return func of the adder outer func, then each returned func has it's own sum var, and therefore it's own state
