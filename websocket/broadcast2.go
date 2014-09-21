package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"code.google.com/p/go.net/websocket"
)

var (
	mux       sync.Mutex
	clientMap = make(map[*websocket.Conn]bool)
)

func addClient(ws *websocket.Conn) {
	mux.Lock()
	defer mux.Unlock()

	clientMap[ws] = true
}

func removeClient(ws *websocket.Conn) {
	mux.Lock()
	defer mux.Unlock()

	delete(clientMap, ws)
}

func broadcastMessage(v interface{}) {
	clErr := make([]*websocket.Conn, 0)

	for cl, _ := range clientMap {
		err := websocket.Message.Send(cl, v)
		if err != nil {
			log.Println(err)
			clErr = append(clErr, cl)
		}
	}

	for _, cl := range clErr {
		removeClient(cl)
	}
}

func echoHandler(ws *websocket.Conn) {
	addClient(ws)
	log.Printf("connected. %s", ws.RemoteAddr().String())

	for {
		var data []byte
		err := websocket.Message.Receive(ws, &data)
		if err != nil {
			log.Println(err)
			removeClient(ws)
			break
		}

		log.Println(string(data))

		broadcastMessage(string(data))
	}

	log.Printf("disconnected. %s", ws.RemoteAddr().String())
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Hello, Gopher %v", time.Now())
	fmt.Fprintf(w, "Hello, Gopher %v", msg)

	broadcastMessage(msg)
}

func main() {
	http.Handle("/echo", websocket.Handler(echoHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/hello", helloHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
