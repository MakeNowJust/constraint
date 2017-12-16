// Package constraint implements '// +build ...' like constraint.
//
// See: https://golang.org/pkg/go/build/#hdr-Build_Constraints
package constraint

// Constraint is an internal representation of constraint.
type Constraint struct {
	conditions []*condition
}

type condition struct {
	trueTags  []string
	falseTags []string
}

// Eval returns a contraint result on given tags.
func (c *Constraint) Eval(tags TagSet) bool {
	for _, cond := range c.conditions {
		if cond.eval(tags) {
			return true
		}
	}

	return false
}

func (cond *condition) eval(tags TagSet) bool {
	for _, tt := range cond.trueTags {
		if !tags.HasTag(tt) {
			return false
		}
	}

	for _, ft := range cond.falseTags {
		if tags.HasTag(ft) {
			return false
		}
	}

	return true
}
