package jsontest

import (
	"encoding/json"
	"fmt"
	// "reflect"
	"testing"
	// "unsafe"
)

type Sample struct {
	A int
	B string
}

func(s *Sample) MarshalJSON() ([]byte, error) {
	fmt.Println("MarshalJSON")
	return []byte(`{"A":1,"B":"Hello"}`), nil
}

func TestMarshal(t *testing.T) {
	var st struct {
				A int
				B string
			}
	st.A = 1
	st.B = "Hello"

	s := Sample{A: 1, B: "Hello"}

	if b, err := json.Marshal(&s); err != nil {
		t.Errorf("json.Marshal Error %s", err)
	} else {
		fmt.Println(string(b))
	}

	// unsafeの意味不明アドレスを突っ込んでpanicのテストは無理だった…。
	// var p unsafe.Pointer
	// p, ok := 1234.(unsafe.Pointer)
	// var v reflect.Value
	
	// v = reflect.NewAt(reflect.TypeOf(st), p)
	// if b, err := json.Marshal(v); err != nil {
	// 	t.Errorf("json.Marshal Error %s", err)
	// } else {
	// 	fmt.Println(string(b))
	// }

	if b, err := json.Marshal(st); err != nil {
		t.Errorf("json.Marshal Error %s", err)
	} else {
		fmt.Println(string(b))
	}

	if b, err := json.Marshal(&st); err != nil {
		t.Errorf("json.Marshal Error %s", err)
	} else {
		fmt.Println(string(b))
	}

}