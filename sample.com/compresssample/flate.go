package main

import (
	"archive/tar"
	"compress/flate"
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"log"
)

func main() {
	var file *os.File
	var err error
	var writer *flate.Writer
	var body []byte

	if file, err = os.Create("output/sample.tar.flate"); err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// flateはレベルを指定する。flate.BestSpeed or flate.BestCompression or DefaultCompression
	if writer, err = flate.NewWriter(file, flate.BestCompression); err != nil {
		log.Fatalln(err)
	}
	defer writer.Close()

	var filepaths = []string {
		"files/b0044482_1413812.jpg",
		"files/dart_flight_school.png",
		"files/golang.txt",
	}

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	for _, filepath := range filepaths {
		if body, err = ioutil.ReadFile(filepath); err != nil {
			log.Fatalln(err)
		}

		if body != nil {
			hdr := &tar.Header {
				Name: path.Base(filepath),
				Size: int64(len(body)),
			}
			if err := tw.WriteHeader(hdr); err != nil {
				println(err)
			}
			if _, err := tw.Write(body); err != nil {
				println(err)
			}
		}
	}

	writer.Write(buf.Bytes())
	writer.Flush()
}