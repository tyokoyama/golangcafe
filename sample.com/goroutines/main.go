package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	// goroutineの使い方（その2）
	go useGoroutine(ch)

	<-ch 					// Channelから結果(1)を受け取る。
	// Output:
	// Goroutineの処理
}

func useGoroutine(ch chan int) {
	fmt.Println("Goroutineの処理")
	ch <- 1 				// Channelに結果(1)を送る。
}