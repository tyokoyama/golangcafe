package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/context"

)

func main() {
	r := mux.NewRouter()


	r.HandleFunc("/", IndexHandler)

	log.Fatal(http.ListenAndServe(":8080", r))

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	for _, param := range strings.Split(r.URL.RawQuery, "&") {
		params := strings.Split(param, "=")
		context.Set(r, params[0], params[1])
	}

	fmt.Fprintf(w, "IndexHandler\n")
	fmt.Fprintf(w, "hoge = %s\n", context.Get(r, "hoge"))
	fmt.Fprintf(w, "fuga = %s\n", context.Get(r, "fuga"))
}
