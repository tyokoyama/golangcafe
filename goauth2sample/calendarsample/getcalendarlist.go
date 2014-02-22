// 事前に以下の事をしておくこと。
// go getコマンドでパッケージを取得しておくこと。
// go get code.google.com/p/goauth2/oauth
// 
// Cloud ConsoleのCredentialsでClient IDを作成しておく。
// Cloud ConsoleのAPIsでAPIをONにしておくこと。
package main

import (
	"code.google.com/p/goauth2/oauth"

    "flag"
    "fmt"
    "io"
    "os"
)

var (
	cachefile = "cache.json"

    // clientID、secret、redirect_urlはDevelopers ConsoleのCredentialsからコピー＆ペーストして下さい。
    clientID = "[your clientID]"
    secret = "[your secret]"
    redirect_url = "[your redirectURIs]"
	scope = "https://www.googleapis.com/auth/calendar"
    // request_urlは使用するAPIのURLを指定して下さい。（この例ではCalendarList）
	request_url = "https://www.googleapis.com/calendar/v3/users/me/calendarList"
    request_token_url = "https://accounts.google.com/o/oauth2/auth"
    auth_token_url = "https://accounts.google.com/o/oauth2/token"
)

func main() {
    flag.Parse()

    // 認証コードを引数で受け取る。
    code := flag.Arg(0)

    config := &oauth.Config{
            ClientId:     clientID,
            ClientSecret: secret,
            RedirectURL:  redirect_url,
            Scope:        scope,
            AuthURL:      request_token_url,
            TokenURL:     auth_token_url,
            TokenCache:   oauth.CacheFile(cachefile),
    }

    transport := &oauth.Transport{Config: config}

    // キャッシュからトークンファイルを取得
    _, err := config.TokenCache.Token()
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

// config.TokenCache.Token()かtransport.Exchange(code)の戻り値で取得したTokenを設定する
// ようだが、無くても動作はしているようにみえる…。高速化？
//    transport.Token = token

    // Calendar APIにアクセス
    r, err := transport.Client().Get(request_url)
    if err != nil {
        fmt.Println("Get: ", err)
        return
    }

    defer r.Body.Close()

    // Write the response to standard output.
    io.Copy(os.Stdout, r.Body)

    // Send final carriage return, just to be neat.
    fmt.Println()

}