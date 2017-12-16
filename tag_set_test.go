package constraint

import (
	"testing"
)

func TestNewTagSet(t *testing.T) {
	tags, err := NewTagSet("")

	if err == nil {
		t.Fatalf("NewTagSet succeeded: %+#v", tags)
	}

	if err != ErrInvalidTag {
		t.Fatalf("NewTagSet returns an unexpected error: %+#v", err)
	}
}
