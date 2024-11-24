package main

import (
	"fmt"
)

func main() {
	m := make(map[string] string)

	m["aa"] = "aa value"
	m["ab"] = "ab value"
	m["a"] = "a value"

	for k, v := range m {
		fmt.Printf("key[%s] = value[%s]\n", k, v)
	}
}