package constraint

import (
	"testing"
)

func TestUnparse(t *testing.T) {
	for _, c := range []*Constraint{
		nil,
		{},
		{conditions: []*condition{{}}},
	} {
		s, err := Unparse(c)
		if err == nil {
			t.Errorf("Unparse(%+#v) succeeded: %+#v", c, s)
			continue
		}

		if err != ErrZeroConstraint {
			t.Errorf("Unparse(%+#v) returns an unexpected error: %+#v", c, err)
		}
	}
}

func TestConstraint_String(t *testing.T) {
	var s string
	defer func() {
		err := recover()
		if err == nil {
			t.Fatalf("Constraint.String succeeded: %q", s)
		}

		if err != ErrZeroConstraint {
			t.Fatalf("Constraint.String panics with an unexpected error: %+#v", err)
		}
	}()

	var c Constraint
	s = c.String()
}
