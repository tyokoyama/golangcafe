package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("sample.so")
	if err != nil {
		fmt.Printf("Open Error[%v]\n", err)
		return 
	}

	s, err := p.Lookup("Print")
	if err != nil {
		fmt.Printf("Lookup Error[%v]\n", err)
		return 
	}

	s.(func(string))("Hello Plugins")
}