package main

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"log"
)

func main() {
	reader, err := zip.OpenReader("output/sample.zip")
	if err != nil {
		log.Fatalln(err)
	}
	defer reader.Close()

	var rc io.ReadCloser
	for _, f := range reader.File {
		rc, err = f.Open()
		if err != nil {
			log.Fatal(err)
		}

		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, rc)
		if err != nil {
			log.Fatal(err)
		}
		if err = ioutil.WriteFile("output/" + f.Name, buf.Bytes(), 0755); err != nil {
			log.Fatal(err)
		}
		rc.Close()
	}
}