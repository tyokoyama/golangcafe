package main

import (
	"fmt"
)

func main() {
	var val any

	val = 1
	fmt.Println(val)
	PrintDetail(val)

	val = "こんにちは世界"
	fmt.Println(val)
	PrintDetail(val)
}

func PrintDetail(v any) {
	switch t := v.(type) {
	case int, int32, int64:
		fmt.Println("int/int32/int64 型", t)
	case string:
		fmt.Println("string 型:", t)
	default:
		fmt.Println("知らない型")
	}
}