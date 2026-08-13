package data

import (
	"time"
)

// All types in the struct are exported so they are visible to encoding/json
// If not exported, they will not be included in the encoding/decoding
type Movie struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"-"` //use the '-' directive to never show in JSON output
	Title     string    `json:"title"`
	Year      int       `json:"year,omitzero"`
	Runtime   int       `json:"runtime,omitzero,string"` //if not wanting to change key name but keep omitzero, can do ",omitzero" -- must keep trailing comma
	// ALSO using string directive to force the JSON value to be a string always
	Genres  []string `json:"genres,omitempty"` //Using `omitempty` to omit empty slices/arrays
	Version int      `json:"version"`
}
