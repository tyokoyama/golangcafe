package main

//import "firstsample"
import (
	"sample.com/firstlib"
	"sample.com/firstlib/random"
)

func main() {
	firstlib.Add(1,1)
	random.RandInt(2)

	Print("Hello, World!")
}