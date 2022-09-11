package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	b := bytes.ReplaceAll([]byte{0x0A, 0x0B, 0x0C}, []byte{0x0B}, []byte{0xff})
	fmt.Printf("% X\n", b)

	re, err := regexp.Compile(`(\d+)年(\d+)月(\d+)日`)
	if err != nil {
		fmt.Errorf("regexp.Compile Error %v", err)
		return
	}
	s := re.ReplaceAllString("1986年01月12日", "${2}/${3} ${1}")
	fmt.Println(s)

	s = re.ReplaceAllLiteralString("1986年01月12日", "${2}/${3} ${1}")
	fmt.Println(s)

	re, err = regexp.Compile(`(^|_)[a-zA-Z]`)
	if err != nil {
		fmt.Errorf("regexp.Compile Error %v", err)
		return
	}

	s = re.ReplaceAllStringFunc("hello_world", func(s string) string {
		return strings.ToUpper(strings.TrimLeft(s, "_"))
	})
	fmt.Println(s)

}