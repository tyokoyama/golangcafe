package main

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"log"
)

func main() {
	var file *os.File
	var err error
	var reader *gzip.Reader

	if file, err = os.Open("output/sample.tar.gz"); err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	if reader, err = gzip.NewReader(file); err != nil {
		log.Fatalln(err)
	}
	defer reader.Close()

	tr := tar.NewReader(reader)

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