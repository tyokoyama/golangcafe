package sample

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	if x := Add(1, 2); x != 3 {
		t.Errorf("Add(1, 2) = %d", x)
	}
	fmt.Println("TestAdd")
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 1)
	}
	fmt.Println("BenchmarkAdd")
}

func ExampleAdd() {
	x := Add(1, 2)
	fmt.Println(x)
	fmt.Println("ExampleAdd")
	// output:
	// 3
	// ExampleAdd
}