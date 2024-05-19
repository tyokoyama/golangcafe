package main

import (
	"flag"
)

func main() {
	max := flag.Int("max", 255, "max value")
	name := flag.String("name", "something", "my name")
	flag.Parse()

	println(*name, *max)
}
