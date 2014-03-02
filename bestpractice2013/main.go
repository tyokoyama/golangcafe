package main

import (
    "bytes"
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
    buf bytes.Buffer
    err  error
}

// Write writes a value to the provided writer in little endian form.
func (w *binWriter) Write(v interface{}) {
    if w.err != nil {
        return
    }
    switch x := v.(type) {
    case string:
        w.Write(int32(len(x)))
        w.Write([]byte(x))
    case int:
        w.Write(int64(x))
    default:
        w.err = binary.Write(&w.buf, binary.LittleEndian, v)
    }
}

// Flush writes any pending values into the writer if no error has occurred.
// If an error has occurred, earlier or with a write by Flush, the error is
// returned.
func (w *binWriter) Flush() (int64, error) {
    if w.err != nil {
        return 0, w.err
    }
    return w.buf.WriteTo(w.w)
}

func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
    bw := &binWriter{w: w}
    bw.Write(g.Name)
    bw.Write(g.AgeYears)

    return bw.Flush()
}

func main() {
	gopher := Gopher{Name: "Takashi Yokoyama", AgeYears: 20}

	size, err := gopher.WriteTo(os.Stdout)

    fmt.Println()
    fmt.Println(size, err)
}