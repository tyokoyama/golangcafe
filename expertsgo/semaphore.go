package main

import (
	"context"
	"log"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	sem := semaphore.NewWeighted(5)

	go do(sem, func() {time.Sleep(1 * time.Second) }, 1)
	go do(sem, func() {time.Sleep(1 * time.Second) }, 2)
	go do(sem, func() {time.Sleep(1 * time.Second) }, 3)

	time.Sleep(5 * time.Second)
}

func do(sem *semaphore.Weighted, f func(), w int64) {
	if err := sem.Acquire(context.Background(), w); err != nil {
		log.Println(err)
		return
	}

	defer sem.Release(w)

	log.Printf("acquired %d", w)

	f()
}