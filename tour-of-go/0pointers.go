package main

import "fmt"

//Go pointers ~ *T = a pointer to a T value
//- it's zero value is nil
//the & operator generates a pointer to its operand (i = 41; p = &i) -> p is now a pointer to i
//the * operator denotes the pointer's underlying value (*p ~ 42) can also set i thru a pointer
// *p = 21 (i is now 21) --> coined as "dereferencing" or "indirecting"
//Go does not have pointer arithmatic

func pointers() {
	i, j := 42, 2701

	p := &i                                      //point to i
	fmt.Println("Read I thru the pointer: ", *p) //read i thru the pointer
	*p = 21                                      //set i thru the pointer
	fmt.Println(i)                               //see the new value of i

	p = &j         // point to j
	*p = *p / 37   //divide j thru the pointer
	fmt.Println(j) //see the new val of j
}
