package oauth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
)

func Hash(key, target string) string {
	// ハッシュ計算
	hashfun := hmac.New(sha1.New, []byte(key))
	hashfun.Write([]byte(target))

	rawsignature := hashfun.Sum(nil)

	// base64エンコード
	base64signature := make([]byte, base64.StdEncoding.EncodedLen(len(rawsignature)))
	base64.StdEncoding.Encode(base64signature, rawsignature)

	return string(rawsignature)
}