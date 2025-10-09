package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	outSlice := make([][]uint8, dy) //top level slice [[uint8, uint8,...],[uint8, uint8,...], ..]
	for y := range outSlice {       //traverse top level
		outSlice[y] = make([]uint8, dx) //each iteration is a slice = []uint ~ [uint8, uint8, uint8]
		for x := range outSlice[y] {
			outSlice[y][x] = uint8(x ^ y)
		}
	}
	return outSlice
}

func main() {
	pic.Show(Pic)
}
