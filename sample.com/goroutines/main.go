package main

import (
	"fmt"
)

type Data struct {
	A int
	B float32
}

func main() {
	ch := make(chan Data)

	// goroutineの使い方（その1）
	go func(ch_local chan Data) {
		fmt.Println("Goroutineの処理")
		ch_local<- Data{A: 1, B: 1.5}				// Channelに結果(1)を送る。
	}(ch)

	st := <-ch 					// Channelから結果(1)を受け取る。

	fmt.Println(st.A, st.B)
	// Output:
	// Goroutineの処理
}
