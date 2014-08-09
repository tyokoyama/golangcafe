package main

import (
	"log"
	"net/url"

	"github.com/gorilla/reverse"
)

func main() {
	var err error

	regexp, err := reverse.CompileRegexp(`/foo/1(\d+)3`)
	if err != nil {
		log.Fatalln(err)
	}

	// url.Valuesの値からURLを生成。
	url, err := regexp.Revert(url.Values{"":{"45678"}})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(url)
}