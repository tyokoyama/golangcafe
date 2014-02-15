package mailsample

import (
	"io/ioutil"
	"net/mail"
	"os"
	"testing"
)

func TestParseAddress(t *testing.T) {
	// ドキュメントにあるメールアドレスでチェック
	address, err := mail.ParseAddress("Barry Gibbs <bg@example.com>")
//	address, err := mail.ParseAddress("bg@example.com")
	if err != nil {
		t.Errorf("%v", err)
	}

	if address.Name != "Barry Gibbs" {
		t.Errorf("address.Name = %s", address.Name)
	}

	if address.Address != "bg@example.com" {
		t.Errorf("address.Address = %s", address.Address)
	}

	t.Logf("address = %s", address)

}

// mail.txtは実行する前に用意して下さい。
// Gmailだと、メールを開いて、右の三角ボタンを押し、メッセージのソースを表示から
// ヘッダ付きのメール本文が表示できます。
// メール本文をファイルにコピーする時に先頭に余分な空白が無いことを確認して下さい。
// 余分な空白があると正しく動作しません。
func TestReadMessage(t *testing.T) {
	var file *os.File
	var err error
	var msg *mail.Message

	if file, err = os.Open("mail.txt"); err != nil {
		t.Errorf("%v", err)
	}
	defer file.Close()

	if msg, err = mail.ReadMessage(file); err != nil {
		t.Errorf("%v", err)
	}

	// ここから以下はメール本文に合わせてテストを書き換えて下さい。
	for key, value := range msg.Header {
		t.Logf("%s, %v", key, value)
	}

	body, err := ioutil.ReadAll(msg.Body)
	t.Logf("%s", string(body))

	if address, err := msg.Header.AddressList("To"); err != nil {
		t.Errorf("%v", err)
	} else {
		t.Logf("%v", address)
	}

	if date, err := msg.Header.Date(); err != nil {
		t.Errorf("%v", err)
	} else {
		t.Logf("%v", date)
	}

	t.Logf("%s", msg.Header.Get("Subject"))
}