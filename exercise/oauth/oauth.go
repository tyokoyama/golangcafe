package oauth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	OAUTH_VERSION    = "1.0"
	SIGNATURE_METHOD = "HMAC-SHA1"

	CALLBACK_PARAM         = "oauth_callback"
	CONSUMER_KEY_PARAM     = "oauth_consumer_key"
	NONCE_PARAM            = "oauth_nonce"
	SESSION_HANDLE_PARAM   = "oauth_session_handle"
	SIGNATURE_METHOD_PARAM = "oauth_signature_method"
	SIGNATURE_PARAM        = "oauth_signature"
	TIMESTAMP_PARAM        = "oauth_timestamp"
	TOKEN_PARAM            = "oauth_token"
	TOKEN_SECRET_PARAM     = "oauth_token_secret"
	VERIFIER_PARAM         = "oauth_verifier"
	VERSION_PARAM          = "oauth_version"

	CONSUMER_KEY      = "jMTASlyTnRNmcLb0qFiw"
	CONSUMER_SECRET   = "edejJagYahmI6JcS3qs3Lh01uJDcrKrNtwa7ICT9o"
	REQUEST_TOKEN_URL = "https://api.twitter.com/oauth/request_token"
	AUTHORIZE_URL     = "https://api.twitter.com/oauth/authorize"
	ACCESS_TOKEN_URL  = "https://api.twitter.com/oauth/access_token"

)

var paramMap map[string]string

// リクエストトークンを作成する。
func CreateSignature() string {
	paramMap = make(map[string]string)

	// requestToken生成用の擬似パラメータ作成
	result := "GET" + "&" + url.QueryEscape(REQUEST_TOKEN_URL) + "&"

	now := time.Now()
	timestamp := now.Unix()
	nonce := rand.New(rand.NewSource(now.UnixNano()))

	nonce_value := strconv.FormatInt(nonce.Int63(), 10)
	// パラメータはソートされている必要がある。
	result += url.QueryEscape(CALLBACK_PARAM+"="+"oob") + url.QueryEscape("&")
	result += url.QueryEscape(CONSUMER_KEY_PARAM+"="+CONSUMER_KEY) + url.QueryEscape("&")
	result += url.QueryEscape(NONCE_PARAM+"="+nonce_value) + url.QueryEscape("&")
	result += url.QueryEscape(SIGNATURE_METHOD_PARAM+"="+SIGNATURE_METHOD) + url.QueryEscape("&")
	result += url.QueryEscape(TIMESTAMP_PARAM+"="+strconv.FormatInt(timestamp, 10)) + url.QueryEscape("&")
	result += url.QueryEscape(VERSION_PARAM + "=" + OAUTH_VERSION)

	paramMap[VERSION_PARAM] = OAUTH_VERSION
	paramMap[SIGNATURE_METHOD_PARAM] = SIGNATURE_METHOD
	paramMap[TIMESTAMP_PARAM] = strconv.FormatInt(timestamp, 10)
	paramMap[NONCE_PARAM] = nonce_value
	paramMap[CONSUMER_KEY_PARAM] = CONSUMER_KEY
	paramMap[CALLBACK_PARAM] = "oob"

	fmt.Println(result)

	// ハッシュ計算
	key := url.QueryEscape(CONSUMER_SECRET) + "&" + url.QueryEscape("")
	fmt.Println(key)
	hashfun := hmac.New(sha1.New, []byte(key))
	hashfun.Write([]byte(result))

	rawsignature := hashfun.Sum(nil)

	// base64エンコード
	base64signature := make([]byte, base64.StdEncoding.EncodedLen(len(rawsignature)))
	base64.StdEncoding.Encode(base64signature, rawsignature)

	return string(base64signature)
}

// リクエストトークン
func RequestProvider(token string) (url.Values, error) {
	var err error
	var req *http.Request
	var res *http.Response

	client := &http.Client{}

	if req, err = http.NewRequest("GET", REQUEST_TOKEN_URL, strings.NewReader("")); err != nil {
		return nil, err
	}

	// パラメータはソートされている必要がある。
	oauthHdr := "OAuth "
	oauthHdr += CALLBACK_PARAM + "=\"" + paramMap[CALLBACK_PARAM] + "\","
	oauthHdr += CONSUMER_KEY_PARAM + "=\"" + paramMap[CONSUMER_KEY_PARAM] + "\","
	oauthHdr += NONCE_PARAM + "=\"" + paramMap[NONCE_PARAM] + "\","
	oauthHdr += SIGNATURE_PARAM + "=\"" + url.QueryEscape(token) + "\","
	oauthHdr += SIGNATURE_METHOD_PARAM + "=\"" + paramMap[SIGNATURE_METHOD_PARAM] + "\","
	oauthHdr += TIMESTAMP_PARAM + "=\"" + paramMap[TIMESTAMP_PARAM] + "\","
	oauthHdr += VERSION_PARAM + "=\"" + paramMap[VERSION_PARAM] + "\""

	req.Header.Add("Authorization", oauthHdr)

	req.Header.Add("Content-Length", "0")

	fmt.Println(req)

	res, err = client.Do(req)

	fmt.Println(res)

	if !(res.StatusCode >= http.StatusOK && res.StatusCode <= http.StatusPartialContent) {
		return nil, errors.New(fmt.Sprintf("http Response Error [%d]", res.StatusCode))
	}

	// レスポンスを解析する。
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	value, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, err
	}

	fmt.Println(value)

	return value, err
}

func GetLoginUrl(requestToken string) string {
	values := make(url.Values)
	values.Set("oauth_token", requestToken)
	return AUTHORIZE_URL + "?" + values.Encode()
}

func GetAccessToken(verifyCode string, values url.Values) (url.Values, error) {
	paramMap = make(map[string]string)

	// AccessToken生成用の擬似パラメータ作成
	result := "GET" + "&" + url.QueryEscape(ACCESS_TOKEN_URL) + "&" + url.QueryEscape(values.Get("oauth_token_secret"))

	now := time.Now()
	timestamp := now.Unix()
	nonce := rand.New(rand.NewSource(now.UnixNano()))

	nonce_value := strconv.FormatInt(nonce.Int63(), 10)
	// パラメータはソートされている必要がある。
	result += url.QueryEscape(CALLBACK_PARAM+"="+"oob") + url.QueryEscape("&")
	result += url.QueryEscape(CONSUMER_KEY_PARAM+"="+CONSUMER_KEY) + url.QueryEscape("&")
	result += url.QueryEscape(NONCE_PARAM+"="+nonce_value) + url.QueryEscape("&")
	result += url.QueryEscape(SIGNATURE_METHOD_PARAM+"="+SIGNATURE_METHOD) + url.QueryEscape("&")
	result += url.QueryEscape(TIMESTAMP_PARAM+"="+strconv.FormatInt(timestamp, 10)) + url.QueryEscape("&")
	// AccessTokenを取得する場合はrequestTokenとverifyCodeを含める必要がある
	result += url.QueryEscape(TOKEN_PARAM+"="+values.Get("oauth_token")) + url.QueryEscape("&")
	result += url.QueryEscape(VERIFIER_PARAM+"="+verifyCode) + url.QueryEscape("&")
	// ----
	result += url.QueryEscape(VERSION_PARAM + "=" + OAUTH_VERSION)

	paramMap[VERSION_PARAM] = OAUTH_VERSION
	paramMap[SIGNATURE_METHOD_PARAM] = SIGNATURE_METHOD
	paramMap[TIMESTAMP_PARAM] = strconv.FormatInt(timestamp, 10)
	paramMap[NONCE_PARAM] = nonce_value
	paramMap[CONSUMER_KEY_PARAM] = CONSUMER_KEY
	paramMap[CALLBACK_PARAM] = "oob"
	paramMap[TOKEN_PARAM] = values.Get("oauth_token")
	paramMap[VERIFIER_PARAM] = verifyCode

	// ハッシュ計算
	key := url.QueryEscape(CONSUMER_SECRET) + "&" + url.QueryEscape(values.Get("oauth_token_secret"))
	fmt.Println(key)
	hashfun := hmac.New(sha1.New, []byte(key))
	hashfun.Write([]byte(result))

	rawsignature := hashfun.Sum(nil)

	// base64エンコード
	base64signature := make([]byte, base64.StdEncoding.EncodedLen(len(rawsignature)))
	base64.StdEncoding.Encode(base64signature, rawsignature)

	var err error
	var req *http.Request
	var res *http.Response

	client := &http.Client{}

	if req, err = http.NewRequest("GET", ACCESS_TOKEN_URL, strings.NewReader("")); err != nil {
		return nil, err
	}

	// パラメータはソートされている必要がある。
	oauthHdr := "OAuth "
	oauthHdr += CALLBACK_PARAM + "=\"" + paramMap[CALLBACK_PARAM] + "\","
	oauthHdr += CONSUMER_KEY_PARAM + "=\"" + paramMap[CONSUMER_KEY_PARAM] + "\","
	oauthHdr += NONCE_PARAM + "=\"" + paramMap[NONCE_PARAM] + "\","
	oauthHdr += SIGNATURE_PARAM + "=\"" + url.QueryEscape(string(base64signature)) + "\","
	oauthHdr += SIGNATURE_METHOD_PARAM + "=\"" + paramMap[SIGNATURE_METHOD_PARAM] + "\","
	oauthHdr += TIMESTAMP_PARAM + "=\"" + paramMap[TIMESTAMP_PARAM] + "\","
	oauthHdr += TOKEN_PARAM + "=\"" + paramMap[TOKEN_PARAM] + "\","
	oauthHdr += VERIFIER_PARAM + "=\"" + paramMap[VERIFIER_PARAM] + "\","
	oauthHdr += VERSION_PARAM + "=\"" + paramMap[VERSION_PARAM] + "\""

	req.Header.Add("Authorization", oauthHdr)

	req.Header.Add("Content-Length", "0")

	fmt.Println(req)

	res, err = client.Do(req)

	fmt.Println(res)

	if !(res.StatusCode >= http.StatusOK && res.StatusCode <= http.StatusPartialContent) {
		return nil, errors.New(fmt.Sprintf("http Response Error [%d]", res.StatusCode))
	}

	// レスポンスを解析する。
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	value, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, err
	}

	fmt.Println(value)

	return value, nil
}

// APIへのリクエスト処理
// TODO: Twitterのstatuses/update固定になっている。
func RequestAPI(accessToken url.Values, method, resource string, param map[string]string) (string, error) {
	paramMap = make(map[string]string)

	// AccessToken生成用の擬似パラメータ作成
	// TODO: Queryパラメータを設定する場合としない場合で処理をわける？
	result := method + "&" + url.QueryEscape(resource) + "&"

	now := time.Now()
	timestamp := now.Unix()
	nonce := rand.New(rand.NewSource(now.UnixNano()))

	nonce_value := strconv.FormatInt(nonce.Int63(), 10)
	// パラメータはソートされている必要がある。
	// result += url.QueryEscape(CALLBACK_PARAM+"="+"oob") + url.QueryEscape("&")
	result += url.QueryEscape(CONSUMER_KEY_PARAM+"="+CONSUMER_KEY) + url.QueryEscape("&")
	result += url.QueryEscape(NONCE_PARAM+"="+nonce_value) + url.QueryEscape("&")
	result += url.QueryEscape(SIGNATURE_METHOD_PARAM+"="+SIGNATURE_METHOD) + url.QueryEscape("&")
//	result += url.QueryEscape(param["status"]) + url.QueryEscape("&")
	result += url.QueryEscape(TIMESTAMP_PARAM+"="+strconv.FormatInt(timestamp, 10)) + url.QueryEscape("&")
	// // AccessTokenを取得する場合はrequestTokenとverifyCodeを含める必要がある
	result += url.QueryEscape(TOKEN_PARAM+"="+accessToken.Get("oauth_token")) + url.QueryEscape("&")
	// result += url.QueryEscape(VERIFIER_PARAM+"="+verifyCode) + url.QueryEscape("&")
	// // ----
	result += url.QueryEscape(VERSION_PARAM + "=" + OAUTH_VERSION)

	// 認証パラメータ以外のパラメータは末尾につける。
	// TODO: スペースのエンコードが"+"になるとまずいので、何か対策が必要。
	// TODO: 本文のデータ内の%XX系の記号は%25XXに変換して、ハッシュ計算する必要がある。
	re := regexp.MustCompile("\\++")
	result += url.QueryEscape("&") + url.QueryEscape("status" + "=" + re.ReplaceAllString(url.QueryEscape(param["status"]), "%20"))

fmt.Println(result)

	paramMap[VERSION_PARAM] = OAUTH_VERSION
	paramMap[SIGNATURE_METHOD_PARAM] = SIGNATURE_METHOD
	paramMap[TIMESTAMP_PARAM] = strconv.FormatInt(timestamp, 10)
	paramMap[NONCE_PARAM] = nonce_value
	paramMap[CONSUMER_KEY_PARAM] = CONSUMER_KEY
	// paramMap[CALLBACK_PARAM] = "oob"
	paramMap[TOKEN_PARAM] = accessToken.Get("oauth_token")
	// paramMap[VERIFIER_PARAM] = verifyCode
//	paramMap["status"] = param["status"]

	// ハッシュ計算
	key := url.QueryEscape(CONSUMER_SECRET) + "&" + url.QueryEscape(accessToken.Get("oauth_token_secret"))
	fmt.Println(key)
	hashfun := hmac.New(sha1.New, []byte(key))
	hashfun.Write([]byte(result))

	rawsignature := hashfun.Sum(nil)

	// base64エンコード
	base64signature := make([]byte, base64.StdEncoding.EncodedLen(len(rawsignature)))
	base64.StdEncoding.Encode(base64signature, rawsignature)

	paramMap[SIGNATURE_PARAM] = string(base64signature)

	// リクエスト
	var err error
	var req *http.Request
	var res *http.Response

	client := &http.Client{}

	// TODO:BODYの方はスペースを%20にする必要がある。
	replus := regexp.MustCompile("\\+")
	body := url.QueryEscape("status") + "=" + url.QueryEscape(param["status"])
	body = replus.ReplaceAllString(body, "%20")

	fmt.Println("body: ", body)

	if req, err = http.NewRequest(method, resource, strings.NewReader(body)); err != nil {
		return "", err
	}

	// パラメータはソートされている必要がある。
	oauthHdr := "OAuth "
	// oauthHdr += CALLBACK_PARAM + "=\"" + paramMap[CALLBACK_PARAM] + "\","
	oauthHdr += CONSUMER_KEY_PARAM + "=\"" + paramMap[CONSUMER_KEY_PARAM] + "\","
	oauthHdr += NONCE_PARAM + "=\"" + paramMap[NONCE_PARAM] + "\","
	oauthHdr += SIGNATURE_PARAM + "=\"" + url.QueryEscape(string(base64signature)) + "\","
	oauthHdr += SIGNATURE_METHOD_PARAM + "=\"" + paramMap[SIGNATURE_METHOD_PARAM] + "\","
//	oauthHdr += "status=\"" + url.QueryEscape(paramMap["status"]) + "\","
	oauthHdr += TIMESTAMP_PARAM + "=\"" + paramMap[TIMESTAMP_PARAM] + "\","
	oauthHdr += TOKEN_PARAM + "=\"" + paramMap[TOKEN_PARAM] + "\","
	// oauthHdr += VERIFIER_PARAM + "=\"" + paramMap[VERIFIER_PARAM] + "\","
	oauthHdr += VERSION_PARAM + "=\"" + paramMap[VERSION_PARAM] + "\""

	req.Header.Add("Authorization", oauthHdr)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.Header.Add("Content-Length", strconv.FormatInt(int64(len(body)), 10))

	fmt.Println(req)

	res, err = client.Do(req)

	fmt.Println(res)

	if !(res.StatusCode >= http.StatusOK && res.StatusCode <= http.StatusPartialContent) {
		return "", errors.New(fmt.Sprintf("http Response Error [%d]", res.StatusCode))
	}

	// レスポンスを解析する。
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(string(resBody))

	return string(resBody), nil
}
