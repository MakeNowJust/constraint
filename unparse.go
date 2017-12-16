package constraint

import (
	"bytes"
)

// String returns a string representation of Constraint.
//
// It is the same as 'Unparse(c)' except to panic on error.
//
// NOTE: it does not keep the order of tags if 'NOT' operator is used.
// For example, String against constraint '!foo,bar' returns "bar,!foo".
// But it does not change the meaning of constraint.
func (c *Constraint) String() string {
	s, err := Unparse(c)
	if err != nil {
		panic(err)
	}

	return s
}

// Unparse returns a string representation of Constraint.
//
// It returns an error if given Constraint is nil or contains zero value
// due to preventing invalid constraint creation.
func Unparse(c *Constraint) (string, error) {
	if c == nil || len(c.conditions) == 0 {
		return "", ErrZeroConstraint
	}

	var buf bytes.Buffer
	for i, cond := range c.conditions {
		if i > 0 {
			buf.WriteRune(' ')
		}

		err := unparseCondition(cond, &buf)
		if err != nil {
			return "", err
		}
	}

	return buf.String(), nil
}

func unparseCondition(cond *condition, buf *bytes.Buffer) error {
	if len(cond.trueTags)+len(cond.falseTags) == 0 {
		return ErrZeroConstraint
	}

	written := false

	for _, tt := range cond.trueTags {
		if written {
			buf.WriteRune(',')
		}
		written = true
		buf.WriteString(tt)
	}

	for _, ft := range cond.falseTags {
		if written {
			buf.WriteRune(',')
		}
		written = true
		buf.WriteRune('!')
		buf.WriteString(ft)
	}

	return nil
}
