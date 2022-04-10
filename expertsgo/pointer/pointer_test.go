package pointer

import (
	"testing"
)

type Value struct {
	content [64]byte
}

//go:noinline
func f(v Value) Value {
	return v
}

//go:noinline
func g(v *Value) *Value {
	return v
}

func Benchmark_Value(b *testing.B) {
	b.ReportAllocs()
	var v Value
	for i := 0; i < b.N; i++ {
		f(v)
	}
}

func Benchmark_Pointer(b *testing.B) {
	b.ReportAllocs()
	var v Value
	for i := 0; i < b.N; i++ {
		g(&v)
	}
}
