package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	var max int

	flag.IntVar(&max, "max", 255, "max value")
	flag.StringVar(&name, "name", "", "my name")
	flag.Parse()

	fmt.Println(max, name)

	for _, arg := range flag.Args() {
		fmt.Println(arg)
	}
}
