package go121sample

import (
	"fmt"
	"testing"
)

func TestAllEvenBuggy(t *testing.T) {
	testCases := []int{1, 2, 4, 6}
	for _, v := range testCases {
		t.Run("sub", func(t *testing.T) {
			t.Parallel()
			fmt.Println(v)
			if v&1 != 0 {
				t.Fatal("odd v", v)
			}
		})
	}
}

func TestAllEven(t *testing.T) {
	testCases := []int{0, 2, 4, 6}
	for _, v := range testCases {
		t.Run("sub", func(t *testing.T) {
			t.Parallel()
			fmt.Println(v)
			if v&1 != 0 {
				t.Fatal("odd v", v)
			}
		})
	}
}
