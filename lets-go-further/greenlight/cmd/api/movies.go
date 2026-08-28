package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ErikTonnesen1/greenlight/internal/data"
	"github.com/ErikTonnesen1/greenlight/internal/data/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	//Declare anonymous struct to define what we expect in a POST request
	var input struct {
		Title   string       `json:"title"`
		Year    int          `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	//Initialize a json.Decoder instance it:
	//reads from request body and uses decode method to decode the body contents into the input struct
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	//Copy the values from the input struct to a new Movie struct
	movie := data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	v := validator.New()

	//Use the Valid() method to see if any of the checks failed. If so, use failedValidateResponse() helper to send a response to the client
	if data.ValidateMove(v, movie); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	movie, err = app.models.Movies.Insert(movie)
	if err != nil {
		app.serveErrorResponse(w, r, err)
		return
	}

	//Create a Location header so client knows where to find newly created resource at
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", movie.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"movie": movie}, headers)
	if err != nil {
		app.serveErrorResponse(w, r, err)
	}
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie, err := app.models.Movies.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serveErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serveErrorResponse(w, r, err)
	}
}
