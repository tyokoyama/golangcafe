package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan struct{}, 3)
	doneCh := make(chan struct{})

	go send(ch, doneCh)
	go receive(ch)
	go receive(ch)

	<-doneCh
}

func send(ch, doneCh chan<- struct{}) {
	t := time.NewTimer(3 * time.Second)

	for {
		select {
		case <-t.C:
			close(ch)
			close(doneCh)
			return
		case ch <- struct{}{}:
		}
	}
}

func receive(ch <-chan struct{}) {
	for {
		select {
		case _, ok := <-ch:
			if !ok {
				return
			}
			log.Println("received")
		}
	}
}