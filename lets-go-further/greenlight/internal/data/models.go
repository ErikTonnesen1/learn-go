package data

import (
	"database/sql"
	"errors"
)

// return custom error when record not found in database
var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Movies MovieModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}
