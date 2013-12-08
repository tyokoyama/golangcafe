package main

import (
	"fmt"
)

type T struct {
        Field1 int32
        Field2 int32
}

type T2 struct {
        X [1<<24]byte
        Field int32
}

func main() {
        var x *T
        p1 := &x.Field1
        p2 := &x.Field2
        var x2 *T2
        p3 := &x2.Field

	fmt.Println(x2.X[0])	
	fmt.Println(p1, p2, p3)
}