package main

import (
	"encoding/binary"
    "fmt"
	"io"
	"os"
)

type Gopher struct {
    Name     string
    AgeYears int
}

type binWriter struct {
    w    io.Writer
    size int64
    err  error
}

func (w *binWriter) Write(v interface{}) {
    if w.err != nil {
        return
    }
    switch v.(type) {
    case string:
        s := v.(string)
        w.Write(int32(len(s)))
        w.Write([]byte(s))
    case int:
        i := v.(int)
        w.Write(int64(i))
    default:
        if w.err = binary.Write(w.w, binary.LittleEndian, v); w.err == nil {
            w.size += int64(binary.Size(v))
        }
    }
}

func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
    bw := &binWriter{w: w}
    bw.Write(g.Name)
    bw.Write(g.AgeYears)

    return bw.size, bw.err
}

func main() {
	gopher := Gopher{Name: "Takashi Yokoyama", AgeYears: 20}

	size, err := gopher.WriteTo(os.Stdout)

    fmt.Println()
    fmt.Println(size, err)
}