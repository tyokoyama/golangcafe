package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt, syscall.SIGHUP)

	tc := time.After(5 * time.Second)

	// シグナルを受信していなければ、関数を抜ける。
	defer signal.Stop(c)
	for {
		select {
		case s := <-c:
			fmt.Printf("Signal Receive: %v\n", s)
			if s == os.Interrupt {
				return
			}
		case <- tc:
			syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
		}
	}
}