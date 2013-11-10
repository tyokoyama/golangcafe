package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go process1(ch1)
	go process2(ch2)

	// 複数Channelを待ち受ける場合（全て受け取るパターン）
	for {
		select {
			case res1 := <-ch1:
				fmt.Printf("process1 Finished[%d]\n", res1)
			case res2 := <-ch2:
				fmt.Printf("process2 Finished[%d]\n", res2)
			case <- time.After(5 * time.Second):
				// timeパッケージを使ったタイムアウト（この方式にすると、goroutineの順番がずれる？）
				fmt.Println("Finish!")
				return
		}
	}

	// Output:
	// process1: [01]・・・が大量に。
	// process2: [01] Finished.・・・が大量に。
	// process1 Finished[1]
	// process2 Finished[2]
	// Finish!
}

func process1(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("process1: [%02d]\n", i)
	}

	ch <- 1
}

func process2(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("process2: [%02d]\n", i)
	}

	ch <- 2
}
