package testcover

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	fmt.Println("-----------------")

	retX := Add(-1, 2)
	if retX != -1 {
		t.Errorf("Param x's check Error\n")
	}

	retY := Add(1, -2)
	if retY != -2 {
		t.Errorf("Param y's check Error\n")
	}

	ret := Add(1, 2)
	if ret != 3 {
		t.Errorf("(Error)1 + 2 = %d\n", ret)
	}
}