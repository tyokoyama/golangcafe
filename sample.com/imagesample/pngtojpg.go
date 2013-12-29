package main

import (
	"image"
	"image/png"
	"image/jpeg"
	"os"
)

func main() {
	var file *os.File
	var outFile *os.File
	var img image.Image
	var err error

	if file, err = os.Open("pkg.png"); err != nil {
		println("Error", err)
		return
	}
	defer file.Close()

	if img, err = png.Decode(file); err != nil {
		println("Error", err)
		return
	}

	if outFile, err = os.Create("out_pkg.jpeg"); err != nil {
		println("Error", err)
		return
	}

	option := &jpeg.Options{Quality: 100}
	if err = jpeg.Encode(outFile, img, option); err != nil {
		println()
		return
	}
	defer outFile.Close()
}