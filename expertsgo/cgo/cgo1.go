package main
/*
#include <math.h>
#cgo LDFLAGS: -lm
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.sqrt(4))
}
