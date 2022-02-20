package main

import (
	"fmt"
)

var fib func(n int) <-chan int

func main() {
	fib = func(n int) <-chan int {
		result := make(chan int)
		go func() {
			defer close(result)
			if n <= 2 {
				result <- 1
				return
			}
			result <- <-fib(n-1) + <-fib(n-2)
		}()
		return result
	}

	fmt.Printf("fib(4) = %d\n", <-fib(4))
}
