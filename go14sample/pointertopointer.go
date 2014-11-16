package main

import (
	"fmt"
)

type T int

func (*T) M() {
	fmt.Println("Called T.M()")
}

func main() {
	var x **T

	x = new(*T)

	(*x).M()

	// y := new(*T)

	// y.M()
}