package httptestsample

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// 本来の処理のダミーその1
var sampleHandler = http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello HTTP Test")
})

func TestNormal(t *testing.T) {
	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()

	// リクエストの送信先はテストサーバのURLへ。
	r, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Error by ioutil.ReadAll(). %v", err)
	}

	if "Hello HTTP Test" != string(data) {
		t.Fatalf("Data Error. %v", string(data))
	}
}

func TestTLS(t *testing.T) {
	ts := httptest.NewTLSServer(sampleHandler)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	c := &http.Client{Transport: tr}

	r, err := c.Get(ts.URL)
	if err != nil {
		t.Fatalf("Error by http.Get(). %v", err)
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Fatalf("Error by ioutil.ReadAll(). %v", err)
	}

	if "Hello HTTP Test" != string(data) {
		t.Fatalf("Data Error. %v", string(data))
	}
}

// 本来の処理のダミーその2
var sample2Handler = http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
	ua := r.Header.Get("User-Agent")

	if !strings.Contains(ua, "Android") {
		http.Error(w, "Not Found.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello HTTP Test")
})

func TestAndroid(t *testing.T) {
	var requests [2]*http.Request
	var err error

	ts := httptest.NewServer(sample2Handler)
	defer ts.Close()

	requests[0], err = http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Errorf("NewRequest[0] Error. %v", err)
	}
	requests[0].Header.Add("User-Agent", "Mozilla/5.0 (iPad; CPU OS 8_1_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) CriOS/39.0.2171.50 Mobile/12B440 Safari/600.1.4")

	requests[1], err = http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Errorf("NewRequest[1] Error. %v", err)
	}
	requests[1].Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 4.4.2; 302KC Build/101.0.2c00) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36")

	c := http.DefaultClient

	for pos, req := range requests {
		r, err := c.Do(req)
		if err != nil {
			t.Fatalf("Error by http.Get(). %v", err)
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Error by ioutil.ReadAll(). %v", err)
		}
		defer r.Body.Close()

		if r.StatusCode == 200 {
			if "Hello HTTP Test" != string(data) {
				t.Fatalf("Data Error. %v", string(data))
			}
		} else {
			if r.StatusCode != 404 {
				t.Fatalf("Status Error %d", r.StatusCode)
			}

			if "Not Found.\n" != string(data) {
				t.Fatalf("Data Error. %v", string(data))
			}

			if pos != 0 {
				t.Fatalf("Request Error. %d", pos)
			}
		}
	}

}

func TestAndroidNoServer(t *testing.T) {
	var requests [2]*http.Request
	var err error

	requests[0], err = http.NewRequest("GET", "/hoge", nil)
	if err != nil {
		t.Errorf("NewRequest[0] Error. %v", err)
	}
	requests[0].Header.Add("User-Agent", "Mozilla/5.0 (iPad; CPU OS 8_1_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) CriOS/39.0.2171.50 Mobile/12B440 Safari/600.1.4")

	requests[1], err = http.NewRequest("GET", "/hoge", nil)
	if err != nil {
		t.Errorf("NewRequest[1] Error. %v", err)
	}
	requests[1].Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 4.4.2; 302KC Build/101.0.2c00) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36")

	for pos, req := range requests {
		r := httptest.NewRecorder()

		sample2Handler(r, req)

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Error by ioutil.ReadAll(). %v", err)
		}

		if r.Code == 200 {
			if "Hello HTTP Test" != string(data) {
				t.Fatalf("Data Error. %v", string(data))
			}
		} else {
			if r.Code != 404 {
				t.Fatalf("Status Error %d", r.Code)
			}

			// httptest.ResponseRecorderは改行コードが付いてしまう
			if "Not Found.\n" != string(data) {
				t.Fatalf("Data Error. %v", string(data))
			}

			if pos != 0 {
				t.Fatalf("Request Error. %d", pos)
			}
		}
	}
}