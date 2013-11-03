package firstlib

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	if x := Add(1, 1); x != 2 {
		// 結果が合わない場合はエラーログを出力して終了
		t.Errorf("1 + 1 = %d\n", x)
//		t.Fatalf("1 + 1 = %d\n", x)
//		t.Skipf("[Error]1 + 1 = %d\n", x)
	}

	if x := Add(1, 2); x != 3 {
		t.Errorf("1 + 2 = %d\n", x)
	}
}

func TestSub(t * testing.T) {
	if x := Sub(2, 1); x != 1 {
		t.Errorf("2 - 1 = %d\n", x)
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 1)
	}

	fmt.Println("BenchmarkAdd")
}

func ExampleAdd() {
	x := Add(1, 1)
	fmt.Println(x)
	// Output:
	// 2
}