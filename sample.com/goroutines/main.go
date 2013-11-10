package main

import (
	"fmt"
)

func main() {
	// goroutineの使い方
	go func() {
		fmt.Println("Goroutineの処理")
	}()

	// Output:
	// 何も出ない。
}