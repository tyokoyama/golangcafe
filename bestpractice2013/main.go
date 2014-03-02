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

func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
    err = binary.Write(w, binary.LittleEndian, int32(len(g.Name)))
    if err == nil {
        size += 4
        var n int
        n, err = w.Write([]byte(g.Name))
        size += int64(n)
        if err == nil {
            err = binary.Write(w, binary.LittleEndian, int64(g.AgeYears))
            if err == nil {
                size += 4
            }
            return
        }
        return
    }
    return
}

func main() {
	gopher := Gopher{Name: "Takashi Yokoyama", AgeYears: 20}

	size, err := gopher.WriteTo(os.Stdout)

    fmt.Println()
    fmt.Println(size, err)
}