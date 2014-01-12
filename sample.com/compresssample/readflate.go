package main

import (
	"archive/tar"
	"bytes"
	"compress/flate"
	"io"
	"io/ioutil"
	"os"
	"log"
)

func main() {
	var file *os.File
	var err error

	if file, err = os.Open("output/sample.tar.flate"); err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	tr := tar.NewReader(flate.NewReader(file))

	var header *tar.Header
	for {
		header, err = tr.Next()
		if err == io.EOF {
			// ファイルの最後
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		buf := new(bytes.Buffer)
		if _, err = io.Copy(buf, tr); err != nil {
			log.Fatalln(err)
		}

		if err = ioutil.WriteFile("output/" + header.Name, buf.Bytes(), 0755); err != nil {
			log.Fatal(err)
		}
	}

}