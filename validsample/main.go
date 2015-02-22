package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"

	"github.com/tyokoyama/golangcafe/validsample/validation"
)

type User struct {
	Name string		`json: "name",valid:required`
	Age int 		`json: "age"`
	Job string	`json: "job", valid:required`
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var target User

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(data, target)
	if err != nil {
		http.Error(w, fmt.Sprintf("data = [%s]", string(data)), http.StatusInternalServerError)
		return
	}

	if err := validation.Valid(target); err != nil {
		http.Error(w, fmt.Sprintf("%s", err.Error()), http.StatusInternalServerError)
	}
}
