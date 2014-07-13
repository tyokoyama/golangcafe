package main

import (
	"code.google.com/p/goauth2/oauth"

	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	// "os"
	// "io"
	"log"
	// "strings"
)

var (
	cachefile = "cache.json"

	client_id = "81306489611-00pbr9h9e513945aoc6datg7rnhrg76c.apps.googleusercontent.com"
	client_secret = "IUOWcRjT_kbbXpzaKlrvb0HQ"
	redirect_uris = "urn:ietf:wg:oauth:2.0:oob"

	scope = "https://www.googleapis.com/auth/gmail.readonly"
	request_token_url = "https://accounts.google.com/o/oauth2/auth"
	request_url = "https://www.googleapis.com/gmail/v1/users/tksyokoyama@gmail.com/messages"
    auth_token_url = "https://accounts.google.com/o/oauth2/token"

    message_get_url = "https://www.googleapis.com/gmail/v1/users/tksyokoyama@gmail.com/messages/"
)

type messages struct {
	Id string `json:"id"`
	ThreadId string `json:"threadId"`
}

type list struct {
	Messages []messages `json:"messages"`
	NextPageToken string `json:"nextPageToken"`
	ResultSizeEstimate uint32 `json:"resultSizeEstimate"`
}

type headers struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type payload struct {
	PartId string `json:"partId"`
	MimeType string `json:"mimeType"`
	Filename string `json:"filename"`
	Headers []headers `json:"headers"`
	Body body `json:"body"`
}

type body struct {
//	AttachmentId string `json:"attachmentId"`
	Data string `json:"data"`
	Size int32 `json:"size"`
}

type thread struct {
	Id string `json:"id"`
	ThreadId string `json:"threadId"`
	LabelIds []string `json:"labelIds"`
	Snippet string `json:"snippet"`
	HistoryId string `json:"historyId"`
	Payload payload `json:"payload"`
	SizeEstimate uint32 `json:"sizeEstimate"`
//	Parts []string `json:"parts"`
}

func main() {
	var l list
	var t thread
	var err error

	flag.Parse()

    // 認証コードを引数で受け取る。
    code := flag.Arg(0)

	config := &oauth.Config{
            ClientId:     client_id,
            ClientSecret: client_secret,
            RedirectURL:  redirect_uris,
            Scope:        scope,
            AuthURL:      request_token_url,
            TokenURL:     auth_token_url,
            TokenCache:   oauth.CacheFile(cachefile),
    }

    transport := &oauth.Transport{Config: config}

    // キャッシュからトークンファイルを取得
    _, err = config.TokenCache.Token()
    if err != nil {
        // キャッシュなし

        // 認証コードなし＝＞ブラウザで認証させるためにURLを出力
        if code == "" {
            url := config.AuthCodeURL("")
            fmt.Println("ブラウザで以下のURLにアクセスし、認証して下さい。")
            fmt.Println(url)
            return
        }

        // 認証トークンを取得する。（取得後、キャッシュへ）
        _, err = transport.Exchange(code)
        if err != nil {
            fmt.Println("Exchange: ", err)
            return
        }

    }

    // Users.messages.listにアクセス
    r, err := transport.Client().Get(request_url)
    if err != nil {
        fmt.Println("Get: ", err)
        return
    }

    defer r.Body.Close()

    buf := bytes.Buffer{}
    if _, err = buf.ReadFrom(r.Body); err != nil {
    	log.Fatalln(err)
    }

    if err = json.Unmarshal(buf.Bytes(), &l); err != nil {
    	fmt.Println(buf.String())
    	log.Fatalln(err)
    }

    // // Write the response to standard output.
    // io.Copy(os.Stdout, r.Body)

    // Users.messages.getにアクセス。（最初の1件目）
	r2, err := transport.Client().Get(message_get_url + l.Messages[0].Id)
	if err != nil {
        fmt.Println("Get: ", err)
        return
	}

    // // Write the response to standard output.
    // io.Copy(os.Stdout, r2.Body)

    buf2 := bytes.Buffer{}
    if _, err = buf2.ReadFrom(r2.Body); err != nil {
    	log.Fatalln(err)
    }

    if err = json.Unmarshal(buf2.Bytes(), &t); err != nil {
    	fmt.Println(buf2.String())
    	log.Fatalln(err)
    }

    // 本文をエンコードする場合はURLEncodingを使う。
    data, encerr := base64.URLEncoding.DecodeString(string(t.Payload.Body.Data))
    if encerr != nil {
        // fmt.Println("base64", t.Payload.Body.Data[190])
        log.Fatalln(encerr)
    }

    fmt.Println(string(data))

//    fmt.Println(t)

    fmt.Println()

}