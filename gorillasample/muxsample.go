package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type MyRoundTripper struct {
}

func (r MyRoundTripper) RoundTrip(req *http.Request) (res *http.Response, err error) {
	fmt.Println("RoundTrip")

	return nil, nil
}

func main() {
	// rt := MyRoundTripper{}
	r := mux.NewRouter()

	// RouterにHandlerを登録する。
	// 正規表現に一致しないURLは全て404になる。
	// 末尾に"/"をつけると、付けないと404になる。
	// Host()を使うとアクセス制限ができる。
	r.HandleFunc("/", IndexHandler).Host("www.sample.com")
	r.HandleFunc("/articles/{category}/", ArticleHandler).Methods("POST")
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)

	r.NotFoundHandler = http.HandlerFunc(NotFound)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "IndexHandler\n")
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "category = %s, id = %s", vars["category"], vars["id"])
}

// NotFoundを返すが、リクエストのログを取得したい場合などに利用できそう。
func NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Println("NotFound")

	http.Error(w, "404 page not found", http.StatusNotFound)
}
