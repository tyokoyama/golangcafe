package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)

	// プロセスが受け取るシグナルを設定する。
	// 受け取ると、Channelにデータが送られてくる。
	signal.Notify(c, os.Interrupt, os.Kill)

	s := <-c
	fmt.Println("Got Signal:", s)
}