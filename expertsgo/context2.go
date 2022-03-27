package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if err := doSomeThingParallel(5); err != nil {
		fmt.Println(err)
	}
}

func doSomeThingParallel(workerNum int) error {
	ctx := context.Background()
	// このcancelは自分の処理をキャンセルするのではなく、Contextのキャンセル処理をする関数のことなのか。
	cancelCtx, cancel := context.WithCancel(ctx)

	defer cancel()

	errCh := make(chan error, workerNum)
	wg := sync.WaitGroup{}
	for i := 0; i < workerNum; i++ {
		i := i
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			if err := doSomeThingWithContext(cancelCtx, num); err != nil {
				cancel()
				errCh <- err
			}
			return
		}(i)
	}

	wg.Wait()

	close(errCh)
	var errs []error
	for err := range errCh {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errs[0]
	}

	return nil
}

func doSomeThingWithContext(ctx context.Context, num int) error {
	select {
	case <- ctx.Done():
		fmt.Println("Allready canceled.")
		return ctx.Err()
	default:
		r := rand.Int() % 2
		if r == 1 {
			// return ctx.Err()
			return fmt.Errorf("Random Error")
		}
	}

	fmt.Println(num)
	return nil
}