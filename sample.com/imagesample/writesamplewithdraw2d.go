package main 

import (
    "code.google.com/p/draw2d/draw2d"
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

	// 描画の前にフォントのディレクトリを設定しておく。
	// フォント名はルールに従って
	draw2d.SetFontFolder("font/")
//	draw2d.SetFontFolder("/System/Library/Fonts/")
    i := image.NewRGBA(img.Bounds())
    gc := draw2d.NewGraphicContext(i)

    // 画像のコピーをしようと思ったら、DrawImage()を呼び出して、元の画像を描画させる？
    gc.DrawImage(img)

    // 任意の線を描画
    gc.MoveTo(60.0, 10.0)
    gc.LineTo(60.0, 20.0)
    gc.MoveTo(65.0, 10.0)
    gc.LineTo(65.0, 20.0)
    gc.MoveTo(58.0, 13.0)
    gc.LineTo(68.0, 13.0)
    gc.MoveTo(58.0, 16.0)
    gc.LineTo(68.0, 16.0)
    gc.Stroke()
    gc.StrokeStringAt("golang", 70, 10)
//    gc.FillStringAt("golang", 70, 10)

	if outFile, err = os.Create("out_write_pkg.png"); err != nil {
		println("Error", err)
		return
	}
	defer outFile.Close()

	if err = png.Encode(outFile, i); err != nil {
		println("Error", err)
		return
	}
}