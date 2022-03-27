package main

import (
	"fmt"
)

type Number interface { 
    ~int | ~int32 | ~int64 | ~float32 | ~float64
}

type NewInt int

func Max[T Number] (x, y T) T {
	if x >= y {
		return x
	}
	return y
}

func main() {
	var x, y NewInt = 1, 2

	max := Max(x, y) // max == NewInt(2)

	fmt.Println(max)
}