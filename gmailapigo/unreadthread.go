package main

import (
	"code.google.com/p/goauth2/oauth"

	"bytes"
	"encoding/json"
	"encoding/base64"
	"flag"
	"fmt"
	"log"

)

var (
	cachefile = "cache.json"

	client_id = "81306489611-00pbr9h9e513945aoc6datg7rnhrg76c.apps.googleusercontent.com"
	client_secret = "IUOWcRjT_kbbXpzaKlrvb0HQ"
	redirect_uris = "urn:ietf:wg:oauth:2.0:oob"

	scope = "https://www.googleapis.com/auth/gmail.modify"
	request_token_url = "https://accounts.google.com/o/oauth2/auth"
	request_url = "https://www.googleapis.com/gmail/v1/users/tksyokoyama@gmail.com/threads"
    auth_token_url = "https://accounts.google.com/o/oauth2/token"

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

type thread2 struct {
	Id string `json:"id"`
	Snippet string `json:"snippet"`
	HistoryId string `json:"historyId"`
}

type response struct {
	Threads []thread2 `json:"threads"`
	NextPageToken string `json:"nextPageToken"`
	ResultSizeEstimate uint32 `json:"resultSizeEstimate"`
}

func main() {
	var res response
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
        	log.Fatalln("Exchange: ", err)
        }

    }

    // Users.messages.listにアクセス
    r, err := transport.Client().Get(request_url + "?q=is:unread")
    if err != nil {
    	log.Fatalln("Get: ", err)
    }

    defer r.Body.Close()

    buf := bytes.Buffer{}
    if _, err = buf.ReadFrom(r.Body); err != nil {
    	log.Fatalln(err)
    }

    if err = json.Unmarshal(buf.Bytes(), &res); err != nil {
    	fmt.Println(buf.String())
    	log.Fatalln(err)
    }

    // fmt.Println(buf.String())

    for _, ths := range res.Threads {
    	var th thread

		r2, err := transport.Client().Get(request_url + "/" + ths.Id)
	    if err != nil {
	    	log.Fatalln("Get: ", err)
	    }

	    buf2 := bytes.Buffer{}
	    if _, err = buf2.ReadFrom(r2.Body); err != nil {
	    	fmt.Println(buf2.String())
	    	log.Fatalln(err)
	    }

	    if err = json.Unmarshal(buf2.Bytes(), &th); err != nil {
	    	fmt.Println(buf2.String())
	    	log.Fatalln(err)
	    }

//	    fmt.Println(buf2.String())

	    for _, header := range th.Payload.Headers {
	    	if header.Name == "Subject" {
	    		fmt.Println(header.Value)

			    // 本文をエンコードする場合はURLEncodingを使う。
			    data, encerr := base64.URLEncoding.DecodeString(string(th.Payload.Body.Data))
			    if encerr != nil {
			        fmt.Println("base64", th.Payload.Body.Data[190])
			        log.Fatalln(encerr)
			    }

			    fmt.Println(string(data))

	    		break
	    	}
	    }

		r2.Body.Close()
    }
}