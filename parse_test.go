package constraint

import (
	"testing"
)

func TestParse(t *testing.T) {
	for _, row := range []struct {
		s   string
		err error
	}{
		{s: "", err: ErrEmptyString},
		{s: "invalid&tag", err: ErrInvalidTag},
	} {
		c, err := Parse(row.s)

		if err == nil {
			t.Errorf("Parse(%q) succeeded: %+#v", row.s, c)
			continue
		}

		if err != row.err {
			t.Errorf("Parse(%q) returns an unexpected error: %+#v", row.s, err)
		}
	}
}
