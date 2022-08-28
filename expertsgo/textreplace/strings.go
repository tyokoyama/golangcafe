package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := strings.Replace("郷に入っては郷に従え", "郷", "Go", 1)
	fmt.Println(s)

	s = strings.Replace("郷に入っては郷に従え", "郷", "Go", -1)
	fmt.Println(s)

	r := strings.NewReplacer("郷", "Go", "入れば", "入っては")
	s = r.Replace("郷に入れば郷に従え")
	fmt.Println(s)

	toUpper := func(r rune) rune {
		if 'a' <= r && r <= 'z' {
			return r - 'a' + 'A'
		}
		return r
	}

	m := strings.Map(toUpper, "Hello, World")
	fmt.Println(m)

	m = strings.Map(unicode.ToUpper, "Hello, World")
	fmt.Println(m)
}
