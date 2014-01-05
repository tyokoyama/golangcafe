package main

import (
	"image"
	"image/png"
	"os"
)


func main() {
	var outFile *os.File
	var err error

	if outFile, err = os.Create("create.png"); err != nil {
		println("Error", err)
		return
	}

	defer outFile.Close()

	rect := image.Rect(0, 0, 100, 100)
	rgba := image.NewRGBA64(rect)

	// #golangとか書きたいけど、とりあえず#だけ
	for i := 0; i < 10; i++ {
		rgba.Set(60, (10 + i), image.Black.At(0, 0))
		rgba.Set(65, (10 + i), image.Black.At(0, 0))
		rgba.Set((58 + i), 13, image.Black.At(0, 0))
		rgba.Set((58 + i), 16, image.Black.At(0, 0))
	}

	outImage := rgba.SubImage(rect)

	if err = png.Encode(outFile, outImage); err != nil {
		println("Error", err)
		return
	}

}