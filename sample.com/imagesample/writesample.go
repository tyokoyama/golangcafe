package main

import (
	"image"
	"image/png"
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

	rect := img.Bounds()
	rgba := image.NewRGBA(rect)
	size := rect.Size()

	// 元画像のコピー
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	// #golangとか書きたいけど、今は時間がない。
	for i := 0; i < 10; i++ {
		rgba.Set(60, (10 + i), image.Black.At(0, 0))
		rgba.Set(65, (10 + i), image.Black.At(0, 0))
		rgba.Set((58 + i), 13, image.Black.At(0, 0))
		rgba.Set((58 + i), 16, image.Black.At(0, 0))
	}

	outRect := image.Rect(0, 0, rect.Max.X, rect.Max.Y)
	outImage := rgba.SubImage(outRect)

	if outFile, err = os.Create("out_write_pkg.png"); err != nil {
		println("Error", err)
		return
	}
	defer outFile.Close()

	if err = png.Encode(outFile, outImage); err != nil {
		println("Error", err)
		return
	}
}