package main
import (
	"fmt"
)
type Stack[T any] []T

// このような書き方ができるようだが、現時点ではコンパイルエラー？
// func ZipWith[S,T,U any](x Stack[T], y Stack[S], func(T, S) U) Stack[U] {
//     // ...
// }

func New[T any]() *Stack[T] {
	v := make(Stack[T], 0)
	return &v
}

func (s *Stack[T]) Push(x T) {
	(*s) = append((*s), x)
}

func (s *Stack[T]) Pop() T {
	v := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return v
}

func main() {
	s := New[string]()
	s.Push("hello")
	s.Push("world")
	fmt.Println(s.Pop()) // world
	fmt.Println(s.Pop()) // hello
}