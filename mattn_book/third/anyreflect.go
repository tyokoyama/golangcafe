package main

import (
	"fmt"
	"reflect"
)

type V int

func main() {
	var val V = 123

	PrintDetail(val)
}

func PrintDetail(v any) {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("int/int32/int64 型", v)
	case reflect.String:
		fmt.Println("string 型:", v)
	default:
		fmt.Println("知らない型")
	}
}