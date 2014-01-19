package ringsample

import (
	"container/ring"
	"testing"
)

type HogeValue int

func(v *HogeValue) Add(value int) {
	*v =  HogeValue(int(*v) + value)
}

func TestDo(t *testing.T) {
	r := ring.New(5)		// 5個の要素

	if r.Len() != 5 {
		t.Errorf("length = %d", r.Len())
	}

	i := 0
	for initialValue := r.Next(); initialValue != r; initialValue = initialValue.Next() {
		value := HogeValue(i)
		initialValue.Value = &value
		i++
	}

	for p := r.Next(); p != r; p = p.Next() {
		t.Logf("value = %d", *p.Value.(*HogeValue))
	}

	r.Do(func(v interface{}) {
		hoge, ok := v.(*HogeValue)
		if ok {
			hoge.Add(1)
		}
	})


	i = 1
	for p := r.Next(); p != r; p = p.Next() {
		check := p.Value.(*HogeValue)
		if int(*check) != i {
			t.Errorf("check = %d, i = %d", *check, i)
		}
		i++
	}
}