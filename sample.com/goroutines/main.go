package main

import (
	"fmt"
)

func main() {
	// goroutineの使い方（その2）
	go useGoroutine()

	// Output:
	// 何も出ない。
}

func useGoroutine() {
	fmt.Println("Goroutineの処理")
}