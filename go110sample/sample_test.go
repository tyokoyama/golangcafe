package go110sample

import (
	"testing"
)

func TestSample(t *testing.T) {
	Sample("sample")

	t.Errorf("Error Sample")
}
