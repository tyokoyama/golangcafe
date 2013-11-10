package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10)				// Channelに長さを指定するとバッファリングできるようになる。

	// goroutineの使い方（その2）を（10個）大量に呼び出す。
	for i := 1; i < 11; i++ {
		go useGoroutine(i, ch)
	}

	for no := range ch {					// Channelから結果(no)を受け取る。
		fmt.Printf("goroutine[%02d] Finished.\n", no)
	}

	// Output:
	// Goroutineの処理[01]・・・が大量に。
	// goroutine[01] Finished.・・・が大量に。
}

func useGoroutine(no int, ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Goroutineの処理[%02d]\n", no)
	}

	ch <- no 				// Channelに結果(no)を送る。
	if no == 10 {			// バッファリング後、これ以上データを送ることがない場合は、明示的にクローズする必要がある。
		close(ch)
	}
}