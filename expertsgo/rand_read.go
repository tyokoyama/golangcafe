package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
)

func main() {
	var b bytes.Buffer
	src := rand.NewSource(0)
	r := rand.New(src)
	io.CopyN(&b, r, 64)

	fmt.Printf("Len: %d\n", b.Len())
	len := b.Len()
	for i := 0; i < len; i++ {
		read, _ := b.ReadByte()
		fmt.Printf("%d: %02X\n", i, read)
	}

	fmt.Printf("\n")
}
