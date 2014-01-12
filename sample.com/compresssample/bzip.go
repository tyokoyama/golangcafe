package main

import (
	"compress/bzip2"
	"io"
	"os"
	"log"
)

func main() {
	var file *os.File
	var err error

	if file, err = os.Open("files/golang.txt.bz2"); err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// bzip2はReaderしかない。
	reader := bzip2.NewReader(file)

	io.Copy(os.Stdout, reader)
}