package oauth

import (
	"fmt"
	"testing"
)

// リクエストトークンの依頼〜ログインURLを取得するまでのテスト
func TestRequestTokenSection(t *testing.T) {
	token := CreateSignature()

	fmt.Println(token)

	requestToken, err := RequestProvider(token)
	if err != nil {
		t.Errorf("%v", err)
	}

	if requestToken == nil {
		t.Errorf("requestToken is Empty")
	}

	fmt.Println(requestToken)

	loginURL := GetLoginUrl(requestToken.Get("oauth_token"))

	fmt.Println(loginURL)
}

// // oauth_tokenを取得する部分のテスト
// func TestAuthTokenSection(t *testing.T) {
// 	token := CreateSignature()

// 	requestToken, err := RequestProvider(token)
// 	if err != nil {
// 		t.Errorf("%v", err)
// 	}

// 	if requestToken == nil {
// 		t.Errorf("requestToken is Empty")
// 	}

// 	if requestToken.Get("oauth_token") == token {
// 		t.Errorf("requestToken no change")
// 	}

// 	loginURL := GetLoginUrl(requestToken.Get("oauth_token"))

// 	fmt.Println(loginURL)
// 	fmt.Println("input verifyCode: ")

// 	verifyCode := ""
// 	fmt.Scanf("%s\n", &verifyCode)

// 	accessToken, err := GetAccessToken(verifyCode, requestToken)
// 	if err != nil {
// 		t.Errorf("%v", err)
// 	}

// 	fmt.Println(accessToken)
// }
