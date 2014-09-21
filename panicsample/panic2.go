package main

import "fmt"

type Hoge struct {
	Fuga string
}

func main() {
	var h *Hoge
	fmt.Println(h.Fuga)		
}
