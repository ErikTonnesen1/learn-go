package dynamarray

import (
	"fmt"
)

type DynamicArray struct {
	size, capacity int
	array          []int
}

func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{
		0,
		capacity,
		make([]int, capacity),
	}
}

func (da *DynamicArray) Get(i int) int {
	return da.array[i]
}

func (da *DynamicArray) Set(i int, n int) {
	if da.array[i] == 0 {
		da.size++
	}
	da.array[i] = n
}

// Pushing to the end of the array means increasing the size by 1
func (da *DynamicArray) Pushback(n int) {
	if n < 0 || n > da.size {
		panic(fmt.Errorf("Invalid index of : %d", n))
	}
	if da.size+1 >= da.capacity {
		da.resize()
	}
	da.array[da.size] = da.array[n]
	da.array[n] = 0
	da.size++
}

// Removes elem from arry
func (da *DynamicArray) Popback() int {
	lastVar := da.array[da.size]
	da.array[da.size] = 0
	da.size--
	return lastVar
}

func (da *DynamicArray) resize() {
	doubleCapacity := da.capacity * 2
	newArray := make([]int, doubleCapacity)
	copied := copy(newArray, da.array)
	if copied != da.size {
		panic(fmt.Errorf("Resize failed: Elements to copy: %d. Elements copied: %d", da.size, copied))
	}
	da.array = newArray
	da.capacity = doubleCapacity
}

func (da *DynamicArray) GetSize() int {
	return da.size
}

func (da *DynamicArray) GetCapacity() int {
	return da.capacity
}
