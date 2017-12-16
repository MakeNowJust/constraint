package constraint

import (
	"strings"
	"unicode"
)

// Parse parses a given string as constraint, then return a Constraint
// which represents it.
func Parse(s string) (*Constraint, error) {
	fs := strings.Fields(s)
	if len(fs) == 0 {
		return nil, ErrEmptyString
	}

	c := new(Constraint)
	for _, f := range fs {
		cond := new(condition)
		for _, t := range strings.Split(f, ",") {
			negated := false
			if len(t) > 0 && t[0] == '!' {
				negated = true
				t = t[1:]
			}

			if !IsValidTag(t) {
				return nil, ErrInvalidTag
			}

			if negated {
				cond.falseTags = append(cond.falseTags, t)
			} else {
				cond.trueTags = append(cond.trueTags, t)
			}
		}

		c.conditions = append(c.conditions, cond)
	}

	return c, nil
}

// IsValidTag checks the given string is valid tag name.
//
// A valid tag consists unicode letters, digits, '_' and '.' only.
// And an empty string is invalid.
func IsValidTag(t string) bool {
	if len(t) == 0 {
		return false
	}

	for _, c := range t {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) && c != '_' && c != '.' {
			return false
		}
	}

	return true
}
