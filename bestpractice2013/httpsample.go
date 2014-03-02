package main

import (
	"log"
	"net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    err := doThis()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        log.Printf("handling %q: %v", r.RequestURI, err)
        return
    }

    err = doThat()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        log.Printf("handling %q: %v", r.RequestURI, err)
        return
    }
}

func main() {
	http.ListenAndServe(":12345", nil)
}

type HogeError struct {
	Message string
}

func (e HogeError) Error() string {
	return e.Message
}

func doThis() error {
//	return HogeError{ Message: "Fire Error doThis"}
	return nil
}

func doThat() error {
	return HogeError{ Message: "Fire Error doThat"}
//	return nil
}