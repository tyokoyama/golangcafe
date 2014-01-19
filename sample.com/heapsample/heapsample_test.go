package heapsample

import (
	"container/heap"
	"testing"
)

func TestAdd(t *testing.T) {
	member := GolangCafe {Name: "sinmetal", Priority: 5, Count: 0, Index: 5}

	heap.Push(golangcafeheap, member)

	if golangcafeheap.Len() < 5 {
		t.Errorf("golangcafeheap length = %d", golangcafeheap.Len())
	}

	n := golangcafeheap.Len()
	for i := 0; i < n; i++ {
		item := golangcafeheap.Pop()
		golangcafe := item.(*GolangCafe)

		t.Logf("Name: %s Priority: %d Count: %d Index: %d",
			 golangcafe.Name, golangcafe.Priority, golangcafe.Count, golangcafe.Index)
	}
}

func TestGet(t *testing.T) {
	if golangcafeheap.Len() <= 0 {
		// Add時に0件になるので…。
		heap.Push(golangcafeheap, GolangCafe {Name: "ttyokoyama", Priority: 1, Count: 13, Index: 2})
		heap.Push(golangcafeheap, GolangCafe {Name: "taknb2nch", Priority: 2, Count: 13, Index: 3})
		heap.Push(golangcafeheap, GolangCafe {Name: "qt_luigi", Priority: 3, Count: 13, Index: 4})
		heap.Push(golangcafeheap, GolangCafe {Name: "tam_x", Priority: 4, Count: 1, Index: 1})
	}

	if golangcafeheap.Len() != 4 {
		t.Errorf("golangcafeheap length = %d", golangcafeheap.Len())
	}

	popItem := heap.Pop(golangcafeheap)

	if golangcafeheap.Len() != 3 {
		t.Errorf("golangcafeheap length = %d", golangcafeheap.Len())
	}

	golangcafe := popItem.(*GolangCafe)
	if golangcafe.Name != "ttyokoyama" {
		t.Errorf("golangcafe.Name = %s", golangcafe.Name)
	}
	t.Logf("Name: %s Priority: %d Count: %d Index: %d",
		 golangcafe.Name, golangcafe.Priority, golangcafe.Count, golangcafe.Index)
}

func TestRemove(t *testing.T) {
	var count int
	if golangcafeheap.Len() <= 0 {
		// Add時に0件になるので…。
		heap.Push(golangcafeheap, GolangCafe {Name: "ttyokoyama", Priority: 1, Count: 13, Index: 2})
		heap.Push(golangcafeheap, GolangCafe {Name: "taknb2nch", Priority: 2, Count: 13, Index: 3})
		heap.Push(golangcafeheap, GolangCafe {Name: "qt_luigi", Priority: 3, Count: 13, Index: 4})
		heap.Push(golangcafeheap, GolangCafe {Name: "tam_x", Priority: 4, Count: 1, Index: 1})
	} else {
		count = golangcafeheap.Len()
	}

	heap.Remove(golangcafeheap, 2)

	if golangcafeheap.Len() != (count - 1) {
		t.Errorf("golangcafeheap.Len() = %d, %d", golangcafeheap.Len(), count)
	}

	n := golangcafeheap.Len()
	for i := 0; i < n; i++ {
		item := golangcafeheap.Pop()
		golangcafe := item.(*GolangCafe)

		t.Logf("Name: %s Priority: %d Count: %d Index: %d",
			 golangcafe.Name, golangcafe.Priority, golangcafe.Count, golangcafe.Index)
	}
}

func TestFix(t *testing.T) {
	if golangcafeheap.Len() <= 0 {
		// Add、Remove時に0件になるので…。
		heap.Push(golangcafeheap, GolangCafe {Name: "ttyokoyama", Priority: 1, Count: 13, Index: 2})
		heap.Push(golangcafeheap, GolangCafe {Name: "taknb2nch", Priority: 2, Count: 13, Index: 3})
		heap.Push(golangcafeheap, GolangCafe {Name: "qt_luigi", Priority: 3, Count: 13, Index: 4})
		heap.Push(golangcafeheap, GolangCafe {Name: "tam_x", Priority: 4, Count: 1, Index: 1})
	}

	heap.Fix(golangcafeheap, 2)

	n := golangcafeheap.Len()
	for i := 0; i < n; i++ {
		item := golangcafeheap.Pop()
		golangcafe := item.(*GolangCafe)

		t.Logf("Name: %s Priority: %d Count: %d Index: %d",
			 golangcafe.Name, golangcafe.Priority, golangcafe.Count, golangcafe.Index)
	}
}