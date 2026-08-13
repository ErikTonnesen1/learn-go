package data

import (
	"fmt"
	"strconv"
)

// Declare a custom Runtime type which has the underlying type int (same as our Movie struct field)
type Runtime int

// implement MarshalJSON so that is satisfies the json.Marshaler interface. Should return the JSON-encoded value for the movie runtime
// Purposefully using value receiver instead of pointer receiver, because a value receiver applies to both values & pointers, whereas pointer receivers work only on pointers
func (r Runtime) MarshalJSON() ([]byte, error) {
	//Generate string of the runtime postfixed with mins
	jsonValue := fmt.Sprintf("%d mins", r)

	//User strconv.Quote to wrap in ""s -- needs to be wrapped in double quotes for correct JSON
	quotedJSONValue := strconv.Quote(jsonValue)

	//convert string to byte string and return it
	return []byte(quotedJSONValue), nil
}
