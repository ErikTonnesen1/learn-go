package main

//a struct is a collection of fields

import "fmt"

type Vertex struct {
	X int
	Y int
}

func structs() {
	vert := Vertex{1, 2}
	p := &vert
	getXFromStruct(vert)
	getXFromPointer(p)
	structLiterals()
}

func getXFromStruct(vert Vertex) {
	vert.X = 4 //struct fields are accessed using a dot (struct.field)
	fmt.Printf("Directly retireved from struct: %d\n", vert.X)
}

//struct fields can be accessed thru a pointer
//to access a field X of a struct when we have a struct pointer, we can write (*p.X) but it is a bit
//cumbersome, so the language allows us to just use p.X and implicitly adds the *

func getXFromPointer(p *Vertex) {
	fmt.Printf("X retrieved from a pointer: %d\n", p.X)
}

// A struct literal denotes a newly allocated struct value by listing the values of its fields.
// You can list a subset of fields by using the Name: __ syntax (order is irrelevant)

func structLiterals() {
	var (
		v1 = Vertex{1, 2}  //has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // has X:0 and Y:0
		p  = &Vertex{1, 2} //has type *Vertex
	)
	fmt.Printf("Regular Vertex: %v\n Pointer: %d\n Other Vertex: %v\n Final vertex: %v\n", v1, p, v2, v3)

}
