package main

import (
	"fmt"
	"github.com/tyokoyama/golangcafe/exercise/oauth"
	"time"
)

func main() {
	token := oauth.CreateSignature()

	requestToken, err := oauth.RequestProvider(token)
	if err != nil {
		panic(err.Error())
	}

	if requestToken == nil {
		panic("requestToken is Empty")
	}

	loginURL := oauth.GetLoginUrl(requestToken.Get("oauth_token"))

	fmt.Println(loginURL)
	fmt.Println("input verifyCode: ")

	verifyCode := ""
	fmt.Scanln(&verifyCode)

	accessToken, err := oauth.GetAccessToken(verifyCode, requestToken)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(accessToken)

	res, err := oauth.RequestAPI(accessToken, "POST", "https://api.twitter.com/1.1/statuses/update.json", map[string]string{"status": "oauth test " + time.Now().String()})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(res)
}
