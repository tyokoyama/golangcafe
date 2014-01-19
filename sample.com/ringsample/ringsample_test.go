package ringsample

import (
	"container/ring"
	"testing"
)

func TestDo(t *testing.T) {
	r := ring.New(5)		// 5個の要素

	if r.Len() != 5 {
		t.Errorf("length = %d", r.Len())
	}

	i := 0
	for initialValue := r.Next(); initialValue != r; initialValue = initialValue.Next() {
		initialValue.Value = i
		i++
	}

	p := r.Next()
	for p != r {
		t.Logf("value = %d", p.Value)
		p = p.Next()
	}

	// addValue := 1
	// r.Do(func (addValue) {
	// 		r.Value = r.Value + addValue
	// 	})


	// i = 1
	// p = r.Next()
	// for p != r {
	// 	if p.Value != i {
	// 		t.Errorf("value = %d, i = %d", p.Value, i)
	// 	}
	// 	p = p.Next()
	// 	i++
	// }
}