package data

import (
	// "encoding/json" //Used in the examples past the struct
	// "fmt"
	"database/sql"
	"errors"
	"github.com/ErikTonnesen1/greenlight/internal/data/validator"
	"github.com/lib/pq"
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

func ValidateMove(v *validator.Validator, movie Movie) {

	//Use Check() to execute validation check
	v.Check(movie.Title != "", "title", "must be provided")
	v.Check(len(movie.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(movie.Year != 0, "year", "must be provided")
	v.Check(movie.Year >= 1888, "year", "must be greater than 1888")
	v.Check(movie.Year <= time.Now().Year(), "year", "must not be in the future")

	v.Check(movie.Runtime != 0, "runtime", "must be provided")
	v.Check(movie.Runtime > 0, "runtime", "must be a positive integer")

	v.Check(movie.Genres != nil, "genres", "must be provided")
	v.Check(len(movie.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(movie.Genres) <= 5, "genres", "must not contain more than 5 genres")

	//Using the Unique helper to check that all values in the movie.Genres slice are unique
	v.Check(validator.Unique(movie.Genres), "genres", "must not contain duplicate values")
}

type MovieModel struct {
	DB *sql.DB
}

func (m MovieModel) Insert(movie Movie) (Movie, error) {
	query := `
	INSERT INTO movies (title, year, runtime, genres)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at, version`

	args := []any{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres)}

	err := m.DB.QueryRow(query, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)

	return movie, err
}

func (m MovieModel) Get(id int) (Movie, error) {

	if id < 1 {
		return Movie{}, ErrRecordNotFound
	}

	query := `
	SELECT id, created_at, title, year, runtime, genres, version
	FROM movies 
	WHERE id = $1`

	movie := Movie{}

	err := m.DB.QueryRow(query, id).Scan(
		&movie.ID,
		&movie.CreatedAt,
		&movie.Title,
		&movie.Year,
		&movie.Runtime,
		pq.Array(&movie.Genres),
		&movie.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return Movie{}, ErrRecordNotFound
		default:
			return Movie{}, err
		}
	}

	return movie, nil
}

//
// func (m MovieModel) Update(movie Movie) (Movie, error) {
//
// 	return nil, nil
// }
//
// func (m MovieModel) Delete(id int) error {
// 	return nil
// }
//
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

// type Movie struct {
// 	ID        int       `json:"id"`
// 	CreatedAt time.Time `json:"-"`
// 	Title     string    `json:"title"`
// 	Year      int       `json:"year,omitzero"`
// 	Runtime   Runtime   `json:"-"` // For this impl, we're voiding the runtime field to -- to never output in the json encoding so that we can use it in the alias further down
// 	Genres    []string  `json:"genres,omitempty"`
// 	Version   int       `json:"version"`
// }
//
// func (m Movie) MarshalJSON() ([]byte, error) {
// 	var runtime string
//
// 	if m.Runtime != 0 {
// 		runtime = fmt.Sprintf("%d mins", m.Runtime)
// 	}
//
// 	//Alias that has the underlying type of Movie
// 	//Due to the way GO handles type definitions, the alias will have all the same fields. But importantly, none of the methods defined on Movie
// 	type MovieAlias Movie
//
// 	//Embed a MovieAlias in the temp struct along with a Runtime field that has the type string and the necessary struct tags
// 	//Important to embed the alias instead of the movie struct so that we don't inherit the MarhsalJSON method that's defined on the Movie struct ~ would cause a infinite loop during encoding
// 	aux := struct {
// 		MovieAlias
// 		Runtime string `json:"runtime,omitzero"`
// 	}{
// 		MovieAlias: MovieAlias(m),
// 		Runtime:    runtime,
// 	}
//
// 	return json.Marshal(aux)
// }

/*
Will result in the following output

{
	"movie": {
		"id": 123,
		"title": "Casablanca",
		"genres": [
			"drama",
			"romance",
			"war"
		],
		"version": 1,
		"runtime": "102 mins"
	}
}

*/
