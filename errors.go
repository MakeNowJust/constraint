package constraint

import (
	"errors"
)

var (
	// ErrEmptyString appears when invoked Parse with an empty string.
	ErrEmptyString = errors.New("an empty string cannot be parsed")

	// ErrZeroConstraint appears when invoked Constraint.String
	// with zero value Constraint.
	ErrZeroConstraint = errors.New("a zero value Constraint is found")

	// ErrInvalidTag appears when a tag name is invalid.
	ErrInvalidTag = errors.New("an invalid tag name is found")
)
