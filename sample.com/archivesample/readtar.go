package main

import (
	"archive/tar"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"log"
)

func main() {
	var file *os.File
	var err error

	if file, err = os.Open("output/sample.tar"); err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// ReaderはClose()はない。
	reader := tar.NewReader(file)

	var header *tar.Header
	for {
		header, err = reader.Next()
		if err == io.EOF {
			// ファイルの最後
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		buf := new(bytes.Buffer)
		if _, err = io.Copy(buf, reader); err != nil {
			log.Fatalln(err)
		}

		if err = ioutil.WriteFile("output/" + header.Name, buf.Bytes(), 0755); err != nil {
			log.Fatal(err)
		}
	}
}