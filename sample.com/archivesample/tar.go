package main

import (
	"archive/tar"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	var file *os.File
	var err error

	if file, err = os.Create("output/sample.tar"); err != nil {
		panic(err)
	}
	defer file.Close()

	// Closeをしないと、展開できなくなる可能性がある。
	// アーカイブユーティリティ（MacのFinderから）だとエラーで展開されない）	
	tw := tar.NewWriter(file)
	defer tw.Close()

	var filepaths = []string {
		"files/b0044482_1413812.jpg",
		"files/dart_flight_school.png",
		"files/golang.txt",
	}

	for _, filepath := range filepaths {
		body := readFile(filepath)
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

}

func readFile(filepath string) (body []byte) {
	var err error
	if body, err = ioutil.ReadFile(filepath); err != nil {
		println(err)
		return
	}
	return
}