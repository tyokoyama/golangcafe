package main

import (
	"io"
	"log"
	"net/http"
	"os"
	// "net/rpc"
	"strings"
)

type HelloArgs struct {
	Who string
}
func main() {
	// HTTP ClientによるRPC呼び出し。
	// jsonのパラメータを自分で作って上げる必要がある。
	client := http.Client{}
	res, err := client.Post("http://localhost:8080/rpc", "application/json", strings.NewReader(`{"method":"HelloService.Say","params":[{"Who":"Test"}], "id":"1"}`))
	if err != nil {
		log.Fatalln(err)
	}

	// 標準のDialHTTPPathはGETのリクエストが投げられるらしい。
	// したがって、標準のrpcパッケージは使えない（？）
	// client, err := rpc.DialHTTPPath("tcp", "localhost:8080", "/rpc")
	// if err != nil {
	// 	log.Println("DialHTTP")
	// 	log.Fatalln(err)
	// }

	// defer client.Close()

	// var reply int

	// args := HelloArgs{Who: "hoge"}
	// err = client.Call("HelloService.Say", args, &reply)
	// if err != nil {
	// 	log.Println("Call")
	// 	log.Fatalln(err)
	// }

	// log.Println(reply)
	io.Copy(os.Stdout, res.Body)
}