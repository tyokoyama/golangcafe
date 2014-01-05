package main

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	var file *os.File
	var err error

	if file, err = os.Create("output/sample.zip"); err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	zw := zip.NewWriter(file)

	var filepaths = []string {
		"files/b0044482_1413812.jpg",
		"files/dart_flight_school.png",
		"files/golang.txt",
	}

	var f io.Writer
	for _, filepath := range filepaths {
		body := readFile(filepath)
		if body != nil {
			f, err = zw.Create(path.Base(filepath))
			if err != nil {
				log.Fatal(err)
			}
			_, err = f.Write(body)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// zip.WriterはClose時にエラーチェックをすること。
	err = zw.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func readFile(filepath string) (body []byte) {
	var err error
	if body, err = ioutil.ReadFile(filepath); err != nil {
		println(err)
		return
	}
	return
}
