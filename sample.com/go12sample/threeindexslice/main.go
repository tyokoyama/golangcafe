package main

import (
	"fmt"
)

func main() {
	var arr [10]int

	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}

	//slice := arr[2:4:7]
	slice := arr[2:4:4]

	var slice2 []int = slice
	for i := 0; i < 5; i++ {
		slice2 = append(slice2, 11 + i)
	}

	fmt.Println(slice, len(slice), cap(slice), arr, slice2, cap(arr))

	slice3 := arr[2:4]
	slice4 := slice3
	for i := 0; i < 10; i++ {
		slice4 = append(slice4, 21 + i)
	}

	fmt.Println(slice3, len(slice3), cap(slice3), arr, slice4, cap(arr))
	fmt.Printf("%v, %[1]T, %[1]T\n", slice4)

	fmt.Printf("%[3]d, %d, %d\n", 1, 2, 3, 4, 5)
}