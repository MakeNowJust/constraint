package constraint

import (
	"errors"
)

var (
	// ErrEmptyString appears when invoked Parse with an empty string.
	ErrEmptyString = errors.New("an empty string cannot be parsed")

	// ErrInvalidTag appears when a tag name is invalid.
	ErrInvalidTag = errors.New("an invalid tag name is found")
)
