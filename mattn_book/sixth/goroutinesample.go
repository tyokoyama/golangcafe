package main

import (
	"fmt"
	"sync"
	// "time"
)

func doSomething(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		fmt.Printf("%d\n", id)
	}
//	time.Sleep(100 * time.Second)
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go doSomething(&wg, i)
	}

	wg.Wait()
}
