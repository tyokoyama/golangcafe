package main

import (
	"fmt"
)

func arrYield(i int) int {
	return i
}

func main() {
	// 配列でも同じことができるのでは？と思ったが・・・
	// こんな事になってしまうのか？
	arr := []func(int) int {
		arrYield,
		arrYield,
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i](i + 1))
	}

	// range over funcだと関数の実装は・・・？
	iter := func(yield func(int) bool) {
		yield(0)
		yield(0)
	}
	for a := range iter {
		fmt.Println(a)
	}
/*
	iter := func(f func(int) int) {
		arrYield(1)
		arrYield(2)
	}
	for index := range iter {
		fmt.Println(index)
	}
*/
}
