package main

//a map maps keys to values
// the zero value of a map is nil. a nil map has no keys, ~nore can keys be added~

//the 'make' function returns a map of the given typek, initialized and ready for use

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func maps() {
	fmt.Println("Make map")
	makeMap()

	fmt.Println("\nMake map literals")
	mapLiterals()

	fmt.Println("\nMake map literals but better")
	mapLiteralsCont()

	fmt.Println("\nMutating maps:")
	mutateMaps()
}

func makeMap() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

//Map literals are like struct literals but the keys are required

func mapLiterals() {
	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

// Map literals cont.
// - if the top-level type is just a type name, you can omit it from the elements of the literal
func mapLiteralsCont() {
	var m = map[string]Vertex{
		"Bell labs": {40.68433, -74.39967},
		"Google":    {37.68433, -122.39967},
	}
	fmt.Println(m)
}

//Mutating Maps

//Insert or update
// - m[key] = elem

//retrieve an element
// - elem = m[key]

//delete an element
// - delete(m, key)

//Test that a key is present with a two-value assignment
// - elem, ok = m[key]
//IF key is in m, ok = true
//IF key is not in m, ok = false

//Note: if elem or ok have not yet been declared, you could use a short declaration form:
// elem, ok := m[key]

func mutateMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
