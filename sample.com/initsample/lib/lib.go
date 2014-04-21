package lib

import (
	"fmt"
)

type Hoge interface {
	Method(int) int
}

func init() {
	fmt.Println("lib package init is called.")
}
