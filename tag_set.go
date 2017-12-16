package constraint

// TagSet is interface of a set of tags.
//
// It is passed to Constraint.Eval to evaluate contraint.
type TagSet interface {
	HasTag(s string) bool
}

// NewTagSet returns a new TagSet with given tags.
//
// NOTE: it is one of the ways to get a TagSet implementation, not the only way.
// Don't forget, that you can implement TagSet for your own type.
func NewTagSet(tags ...string) (TagSet, error) {
	set := make(tagSet)
	for _, t := range tags {
		if IsValidTag(t) {
			set[t] = struct{}{}
		} else {
			return nil, ErrInvalidTag
		}
	}
	return set, nil
}

// tagSet is internal type for NewTagSet.
type tagSet map[string]struct{}

// tagSet should implement TagSet interface.
var _ TagSet = make(tagSet)

// HasTag checks whether tags has the tag named s or not.
func (tags tagSet) HasTag(s string) bool {
	_, ok := tags[s]
	return ok
}
