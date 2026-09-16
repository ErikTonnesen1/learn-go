package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Declare a custom Runtime type which has the underlying type int (same as our Movie struct field)
var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

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

// Implements the Unmarshaler interface. Because UnmarshalJSON() needs to modify the receiver (our Runtime type), we must use a pointer receiver
// for this to work correctly
// Otherwise, we'll only be modifying a copy
func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	//We expect that the incoming JSON value will be a string in the format "<runtime> mins"
	//So we must first remove the ("")'s. If we cannot do that, we throw the ErrInvalidRuntimeFormat error
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.Atoi(parts[0])
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)

	return nil
}
