package main

import (
	"fmt"
)

func main() {
	Print123()
}

func Print123() {
	var prints []func()
	for i := 1; i <= 3; i++ {
		prints = append(prints, func() { fmt.Println(i) })
	}
	for _, print := range prints {
		print()
	}
}
