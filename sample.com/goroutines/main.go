package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	// goroutineの使い方（その1）
	go func(chan int) {
		fmt.Println("Goroutineの処理")
		ch<- 1				// Channelに結果(1)を送る。
	}(ch)

	<-ch 					// Channelから結果(1)を受け取る。
	// Output:
	// Goroutineの処理
}
