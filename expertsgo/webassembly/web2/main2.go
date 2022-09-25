package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	window := js.Global()

	document := window.Get("document")

	body := document.Get("body")

	btn := document.Call("createElement", "button")

	btn.Set("textContent", "click me!")

	btn.Call("addEventListener", "click", js.FuncOf(func(js.Value, []js.Value) interface{} {
		fmt.Println("Hello, WebAssembly!")
		return nil
	}))
	body.Call("appendChild", btn)

	select {}
}

func newUint8Array(size int) js.Value {
	ua := js.Global().Get("Uint8Array")
	return ua.New(size)
}