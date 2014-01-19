package listsample

import (
	"container/list"
	"testing"
)

func TestPushBack(t *testing.T) {
	l := list.New()

	for i := 0; i < 5; i++ {
		l.PushBack(i)		
	}

	if l.Len() != 5 {
		t.Errorf("length = %d", l.Len())
	}

	e := l.Front();				// 先頭の要素を取得
	for i := 0; i < 5; i++ {

		if e.Value != i {
			t.Errorf("e.Value = %d, i = %d", e.Value, i)
		}

		e = e.Next()
	}
}

func TestPushFlont(t *testing.T) {
	l := list.New()

	for i := 0; i < 5; i++ {
		l.PushFront(i)		
	}

	if l.Len() != 5 {
		t.Errorf("length = %d", l.Len())
	}

	e := l.Front();
	for i := 4; i <= 0; i-- {

		if e.Value != i {
			t.Errorf("e.Value = %d, i = %d", e.Value, i)
		}

		e = e.Next()
	}

}

func TestInsertAfter(t *testing.T) {
	var e *list.Element
	l := list.New()

	for i := 0; i < 5; i++ {
		if i == 3 {
			e = l.PushFront(i)
		} else {
			l.PushFront(i)		
		}
	}

	l.InsertAfter(5, e)

	if l.Len() != 6 {
		t.Errorf("length = %d", l.Len())
	}

	e = l.Front()
	for i := 0; i < 6; i++ {
		t.Logf("e = %d", e.Value)

		e = e.Next()
	}
}

func TestInsertBefore(t *testing.T) {
	var e *list.Element
	l := list.New()

	for i := 0; i < 5; i++ {
		if i == 3 {
			e = l.PushFront(i)
		} else {
			l.PushFront(i)		
		}
	}

	l.InsertBefore(5, e)

	if l.Len() != 6 {
		t.Errorf("length = %d", l.Len())
	}

	e = l.Front()
	for i := 0; i < 6; i++ {
		t.Logf("e = %d", e.Value)

		e = e.Next()
	}
}