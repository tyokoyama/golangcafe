package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type T struct {
		A int
		B string
		C []int
		D int
	}

	var v T
	fmt.Printf("v    address = %p\n", &v)
	fmt.Printf("v.A  address = %p\n", &v.A)
	fmt.Printf("v.B  address = %p\n", &v.B)
	fmt.Printf("v.C  address = %p\n", &v.C)
	fmt.Printf("A size = %d\n", uintptr(unsafe.Pointer(&v.B))-uintptr(unsafe.Pointer(&v.A)))
	fmt.Printf("B size = %d\n", uintptr(unsafe.Pointer(&v.C))-uintptr(unsafe.Pointer(&v.B)))
	fmt.Printf("C size = %d\n", uintptr(unsafe.Pointer(&v.D))-uintptr(unsafe.Pointer(&v.C)))
	
}