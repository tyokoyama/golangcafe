package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("mypic.png")
	// file, err := os.Open("22238084004_171cf4e34c_o.jpeg")
	// file, err := os.Open("readpng.go")
	if err != nil {
		fmt.Printf("File Open Error[%v].\n", err)
		os.Exit(1)
	}

	result, err := IsPNG(file)
	if err != nil {
		fmt.Printf("File is not PNG[%v].\n", err)
		file.Close()
		os.Exit(1)
	}

	file.Close()
	if result {
		fmt.Printf("File is PNG.\n")
	} else {
		fmt.Printf("File is not PNG.\n")
	}
}

func IsPNG(r io.Reader) (bool, error) {
	magicnum := []byte{137, 80, 78, 71}
	buf := make([]byte, len(magicnum))
	_, err := io.ReadAtLeast(r, buf, len(buf))
	if err != nil {
		return false, err
	}
	return bytes.Equal(magicnum, buf), nil
}