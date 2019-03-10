package main

import (
	"fmt"
	"reflect"
)

type T1 struct {
    t2
}

type t2 struct {
    X int
    u int
}

func main() {
	var t1 T1
	v1 := reflect.ValueOf(&t1).Elem()
	v2 := v1.Field(0)
	X := v2.Field(0)
	u := v2.Field(1)

	fmt.Printf("%v, %v\n", v1, v1.CanSet())
	fmt.Printf("%v, %v, %v\n", v1, v1.Type().Field(0).PkgPath, v1.Type().Field(0).Anonymous)
	fmt.Printf("%v, %v\n", v2, v2.CanSet())
	fmt.Printf("%v, %v\n", X, X.CanSet())
	fmt.Printf("%v, %v\n", u, u.CanSet())

}