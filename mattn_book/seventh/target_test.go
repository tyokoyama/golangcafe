package seventh

import (
	"testing"
)

func TestStringDistance(t *testing.T) {
	got := StringDistance("foo", "foh")
	want := 2
	if got != want {
		t.Fatalf("expected: %v, got %v:", want, got)
	}
}
