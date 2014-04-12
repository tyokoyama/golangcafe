package oauth

import (
	"math/rand"
	"net/url"
	"strconv"
	"time"
)

type Parameter struct {
	Callback        string // コールバック先
	ConsumerKey     string
	ConsumerSecret  string
	Nonce           string
	SignatureMethod string
	Timestamp       string
	Version         string
}

// Create Parameter object.
func newParameter(Callback, ConsumerKey, ConsumerSecret string) Parameter {
	now := time.Now()
	timestamp := now.Unix()
	nonce := rand.New(rand.NewSource(now.UnixNano()))

	nonce_value := strconv.FormatInt(nonce.Int63(), 10)

	return Parameter{
		Callback:        Callback,
		ConsumerKey:     ConsumerKey,
		ConsumerSecret:  ConsumerSecret,
		Nonce:           nonce_value,
		SignatureMethod: "HMAC-SHA1",
		Timestamp:       strconv.FormatInt(timestamp, 10),
		Version:         "1.0",
	}
}

// update timestamp and nonce.
func (p *Parameter) refresh() {
	now := time.Now()
	timestamp := now.Unix()
	nonce := rand.New(rand.NewSource(now.UnixNano()))

	p.Timestamp = strconv.FormatInt(timestamp, 10)
	p.Nonce = strconv.FormatInt(nonce.Int63(), 10)

	println(p.Timestamp, p.Nonce)
}

// Get Parameter string for RequestToken.
func (p Parameter) getRequestTokenParam() string {
	// パラメータはソートされている必要がある。
	result := ""
	result += url.QueryEscape(CALLBACK_PARAM+"="+p.Callback) + url.QueryEscape("&")
	result += url.QueryEscape(CONSUMER_KEY_PARAM+"="+p.ConsumerKey) + url.QueryEscape("&")
	result += url.QueryEscape(NONCE_PARAM+"="+p.Nonce) + url.QueryEscape("&")
	result += url.QueryEscape(SIGNATURE_METHOD_PARAM+"="+p.SignatureMethod) + url.QueryEscape("&")
	result += url.QueryEscape(TIMESTAMP_PARAM+"="+p.Timestamp) + url.QueryEscape("&")
	result += url.QueryEscape(VERSION_PARAM + "=" + p.Version)

	return result
}
