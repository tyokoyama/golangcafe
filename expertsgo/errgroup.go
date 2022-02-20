package main

import (
	"errors"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	var eg errgroup.Group

	for i := 0; i < 10; i++ {
		n := i
		eg.Go(func() error {
			return do(n)
		})
	}

	if err := eg.Wait(); err != nil {
		log.Printf("err = %v", err)
	}
}

func do(n int) error {
	if n % 2 == 0 {
		return errors.New("err")
	}

	time.Sleep(1 * time.Second)
	log.Printf("%d called", n)

	return nil
}