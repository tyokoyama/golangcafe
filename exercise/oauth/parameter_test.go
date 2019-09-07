package oauth

import (
	"testing"
	"time"
)

// const (
// 	OAUTH_VERSION    = "1.0"
// 	SIGNATURE_METHOD = "HMAC-SHA1"

// 	CALLBACK_PARAM         = "oauth_callback"
// 	CONSUMER_KEY_PARAM     = "oauth_consumer_key"
// 	NONCE_PARAM            = "oauth_nonce"
// 	SESSION_HANDLE_PARAM   = "oauth_session_handle"
// 	SIGNATURE_METHOD_PARAM = "oauth_signature_method"
// 	SIGNATURE_PARAM        = "oauth_signature"
// 	TIMESTAMP_PARAM        = "oauth_timestamp"
// 	TOKEN_PARAM            = "oauth_token"
// 	TOKEN_SECRET_PARAM     = "oauth_token_secret"
// 	VERIFIER_PARAM         = "oauth_verifier"
// 	VERSION_PARAM          = "oauth_version"

// 	CONSUMER_KEY      = "7E7NkNRyseZTWxGSkNHQ"
// 	CONSUMER_SECRET   = "qsZecZRe391NDfzY4F24BzQnoWNQhbMyzWqe6SyQ"
// 	REQUEST_TOKEN_URL = "https://api.twitter.com/oauth/request_token"
// 	AUTHORIZE_URL     = "https://api.twitter.com/oauth/authorize"
// 	ACCESS_TOKEN_URL  = "https://api.twitter.com/oauth/access_token"

// )

func TestNewParameter(t *testing.T) {
	param := newParameter("oob", CONSUMER_KEY, CONSUMER_SECRET)

	if param.Callback != "oob" {
		t.Errorf("callback is not 'oob'. [%s]", param.Callback)
	}

	if param.ConsumerKey != CONSUMER_KEY {
		t.Errorf("ConsumerKey is not %s. [%s]", CONSUMER_KEY, param.ConsumerKey)
	}

	if param.ConsumerSecret != CONSUMER_SECRET {
		t.Errorf("ConsumerSecret is not %s. [%s]", CONSUMER_SECRET, param.ConsumerSecret)
	}

	if param.SignatureMethod != SIGNATURE_METHOD {
		t.Errorf("SignatureMethod is not %s. [%s]", SIGNATURE_METHOD, param.SignatureMethod)
	}

	if param.Version != OAUTH_VERSION {
		t.Errorf("Version is not %s. [%s]", OAUTH_VERSION, param.Version)
	}
}

func TestGetRequestTokenParam(t *testing.T) {
	param := newParameter("oob", CONSUMER_KEY, CONSUMER_SECRET)

	// 仮のタイムスタンプとNonceでテスト
	param.Timestamp = "1396176334"
	param.Nonce = "5429236143078588817"

	result := "oauth_callback%3Doob%26oauth_consumer_key%3D7E7NkNRyseZTWxGSkNHQ%26oauth_nonce%3D5429236143078588817%26oauth_signature_method%3DHMAC-SHA1%26oauth_timestamp%3D1396176334%26oauth_version%3D1.0"

	if param.getRequestTokenParam() != result {
		t.Errorf("RequestParam Error %s", param.getRequestTokenParam())
	}
}

func TestRefresh(t *testing.T) {
	param := newParameter("oob", CONSUMER_KEY, CONSUMER_SECRET)
	param2 := newParameter("oob", CONSUMER_KEY, CONSUMER_SECRET)

	param2.Timestamp = param.Timestamp
	param2.Nonce = param.Nonce
	time.Sleep(100 * time.Millisecond)
	param2.refresh()

	if param.Timestamp == param2.Timestamp {
		t.Errorf("param.Timestamp not refreshed.")
	}

	if param.Nonce == param2.Nonce {
		t.Errorf("param.Nonce not refreshed.")
	}
}
