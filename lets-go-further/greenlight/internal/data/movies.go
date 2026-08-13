package data

import (
	"encoding/json"
	"fmt"
	"time"
)

// All types in the struct are exported so they are visible to encoding/json
// If not exported, they will not be included in the encoding/decoding
type Movie struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"-"` //use the '-' directive to never show in JSON output
	Title     string    `json:"title"`
	Year      int       `json:"year,omitzero"`
	Runtime   Runtime   `json:"runtime,omitzero"` //if not wanting to change key name but keep omitzero, can do ",omitzero" -- must keep trailing comma
	Genres    []string  `json:"genres,omitempty"` //Using `omitempty` to omit empty slices/arrays
	Version   int       `json:"version"`
}

/*


Alternative to using a custom Runtime type, with a MarshalJSON() method, we could just write a MarshalJSON() method for our movie struct

*/

/*
func (m Movie) MarshalJSON() ([]byte, error) {
	var runtime string

	if m.Runtime != 0 {
		runtime = fmt.Sprintf("%d mins", m.Runtime)
	}

	aux := struct {
		ID      int      `json:"id"`
		Title   string   `json:"title"`
		Year    int      `json:"year,omitzero"`
		Runtime string   `json:"runtime,omitzero"`
		Genres  []string `json:"genres,omitzero"`
		Version int      `json:"version"`
	}{
		ID:      m.ID,
		Title:   m.Title,
		Year:    m.Year,
		Runtime: runtime,
		Genres:  m.Genres,
		Version: m.Version,
	}

	return json.Marshal(aux)
}
*/

/*

Although you could say that the above code is verbose and repetitive, therefore you could also Embed and alias



*/
