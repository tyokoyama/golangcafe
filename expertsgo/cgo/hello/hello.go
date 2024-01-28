package hello

//extern void hello();
import "C"
import "fmt"

//export goHello
func goHello() {
	fmt.Println("hello")
}

func Hello() {
	C.hello()
}