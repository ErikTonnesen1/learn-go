package main

import (
	"fmt"
	"math"
)

//Go does NOT have classes.....Hooray!!
// - you can instead define methods on ~types~

//Method => a function with a special ~receiver~ argument
// - a receiver appears in it's own argument list between 'func' and the method name
// i.e. Abs() has a ~receiver~ of type Vertex

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Methods are just functions, here Abs is written without a receiver arg and still has the same functionality
func Abs2(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//You declare methods on types other than structs, i.e we can use a Float type
// NOTE: You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

type MyFloat float64

func (f MyFloat) Abs3() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{4, 3}
	fmt.Println(v.Abs())
	fmt.Println(Abs2(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs3())
}
