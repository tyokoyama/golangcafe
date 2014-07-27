package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

)

func main() {
	r := mux.NewRouter()


	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/session", SessionHandler)

	log.Fatal(http.ListenAndServe(":8080", r))

}

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "mysession")

	session.Values["hoge"] = "1"
	session.Values[2] = 2

	session.Save(r, w)

	fmt.Fprintf(w, "IndexHandler\n")
}

func SessionHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "mysession")

	fmt.Println(session)

	hoge := session.Values["hoge"]
	val := session.Values[2]

	session.Values["hoge"] = "session"
	session.Values[2] = 5

	session.Save(r, w)

	fmt.Fprintf(w, "SessionHandler [%s] [%d]\n", hoge, val)

}