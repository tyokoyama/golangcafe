package main

import (
	"flag"
	"log"
	"net/smtp"
)

func main() {
	// target = 送信先のアドレス
	// myAddress = メールを送信するアカウントのメールアドレス
	target := "target address"
	myAddress := "your gmail address"

	var to []string = []string{target}

	msg := "To: " + to[0] + "\r\nSubject: " + "Mail Sample by Golang\r\n\r\n" + "This mail is mailsample by golang."

	flag.Parse()

	arg := flag.Arg(0)
	log.Printf("arg = %s", arg)

	// Gmailのsmtpを使うためのパスワードを生成するには、以下を参照。
	// （※アカウントのパスワードではない）
	// https://support.google.com/accounts/answer/185833
	auth := smtp.PlainAuth("", myAddress, arg, "smtp.gmail.com")
	// 465(SSL)、587(TLS)、25(SSL)
	err := smtp.SendMail("smtp.gmail.com:587", auth, myAddress, to, []byte(msg))
	if err != nil {
		log.Fatalln(err)
	}
}