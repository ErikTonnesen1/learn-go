package dynamarray

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

// end of array = n
// array[n] = 0
// pushback results in size of underlying array reaching capacity? da.capacity = doubled
// n < 0 ? or > da.capacity ? --> da.capacity = doubled
func TestPushbackWorks(t *testing.T) {
	da := NewDynamicArray(5)
	for i, _ := range da.array {
		da.Set(i, rand.Intn(10))
	}
	fmt.Printf("Array before Pushback: %#v\n With Size: %d, and Capacity: %d\n", da.array, da.size, da.capacity)
	nValue := da.Get(0)
	fmt.Printf("nValue: %d\n", nValue)
	da.Pushback(0)
	fmt.Println("Size and capacity after pushback: ", da.size, " ", da.capacity)
	lastElem := da.Get(da.size - 1)

	fmt.Printf("Array after Pushback: %#v\n", da.array)

	assert.EqualValues(t, nValue, lastElem)
}
