package main

import (
	"fmt"
	"slices"
)

func main() {

	testCase1 := [][]int{{1, 2}, {3, 5}, {2, 2}}
	fmt.Println(maxAverageRatio(testCase1, 2))

}

// find lowest average amongst classes, add 1 extraStudent to that average, re-evaluate and repeat
func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	var avgs []float64 = make([]float64, 3)
	for i, v := range classes {
		avgs[i] = float64(v[0]) / float64(v[1])
	}
	fmt.Println(avgs)
	for range extraStudents {
		lowestAvg := slices.Min(avgs)
		classIndex := slices.Index(avgs, lowestAvg)
		fmt.Println("Index: " + classIndex)
		classes[classIndex][0]++
		classes[classIndex][1]++
		avgs[classIndex] = float64(classes[classIndex][0] / classes[classIndex][1])
		fmt.Println(avgs)
	}

	avgSum := float64(0)
	for _, v := range avgs {
		avgSum += v
	}
	return float64(avgSum / float64(len(classes)))
}
