package heapsample

import (
	"container/heap"
)

type GolangCafeHeap []*GolangCafe
var golangcafeheap *GolangCafeHeap 	// Heapのデータ

// sort.InterfaceのCompositeなので実装しなければいけない。
func (item GolangCafeHeap) Len() int {return len(item)}
func (item GolangCafeHeap) Less(i, j int) bool {
	// 比較条件を記述（今回は、Priorityが小さいかどうかにする。）
	return item[i].Priority < item[j].Priority
}
func (item GolangCafeHeap) Swap(i, j int) {
	priority := item[i].Priority

	// 要素の入れ替え
	item[i], item[j] = item[j], item[i]

	item[j].Priority = item[i].Priority
	item[i].Priority = priority
}

func (item *GolangCafeHeap) Push(x interface{}) {
	addItem := x.(GolangCafe)
	*item = append(*item, &addItem)
}

func (item *GolangCafeHeap) Pop() interface{} {
	old := *item
	popItem := old[len(old) - 1]
	*item = old[0 : (len(old) - 1)]
	return popItem
}

type GolangCafe struct {
	Name string				// 参加者の名前
	Priority int			// 優先順位
	Count int				// 参加回数
	Index int				// ヒープ内のindex
}

func init() {
	golangcafeheap = &GolangCafeHeap{&GolangCafe {Name: "ttyokoyama", Priority: 1, Count: 13, Index: 2},
			  	&GolangCafe {Name: "taknb2nch", Priority: 2, Count: 13, Index: 3},
			  	&GolangCafe {Name: "qt_luigi", Priority: 3, Count: 13, Index: 4},
				&GolangCafe {Name: "tam_x", Priority: 4, Count: 1, Index: 1},
			}
	heap.Init(golangcafeheap)
}
