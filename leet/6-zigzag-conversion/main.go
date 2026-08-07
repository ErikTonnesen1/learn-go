package main

import (
	"internal/godebug"
	"strings"
)

//conditions
// - length of word
// - amount of rows

//"hellosuperworld" , 3
//[h, 0, o, 0, e, 0, r]
//[e, l, s, p, r, o, l]
//[l, 0, u, 0, w, 0, d]

//[h,l,o,u,e,w,r,d]
//[e,l,s,p,r,o,l,0]

//[row, column]
//[0,0]
//[1,0]
//[2,0]
//[1,1]
//[0,2]
//[1,2]
//[2,2]
//[1,3]
//[0,4]
//[1,4]
//[2,4]

func convert(s string, numRows int) (finalAnswer string) {
	if numRows == 1 {
		return s
	}

	//initial state
	rows := make([]string, numRows)
	currRow := 0
	goingDown := true

	for _, v := range s {
		//action
		rows[currRow] += string(v)

		//state logic
		if currRow == 0 || currRow == numRows-1 {
			goingDown = !goingDown
		}

		//update state
		if goingDown {
			currRow++
		} else {
			currRow--
		}
	}

	//create answer

	return finalAnswer
}
