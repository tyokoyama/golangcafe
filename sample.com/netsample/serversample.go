package main

import (
	"encoding/binary"
	"log"
	"net"
)

func main() {
	// Listenはtcp/tcp4/tcp6/unix/unixpacketでなければエラーになる。
	// エラーを繰り返すと、指定したアドレス名のファイルができる？（for MacOSX）
	listener, err := net.Listen("tcp", "localhost:22000")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept: %v\n", err)
			continue
		}

		go receiveGoroutine(conn)
	}
}

func receiveGoroutine(conn net.Conn) {
	var count uint32 = 1

	for i := 0; i < 2; i++ {
		var err error

		log.Printf("Send To Client %d\n", count)
		err = binary.Write(conn, binary.LittleEndian, count)
		if err != nil {
			log.Printf("Send: %v\n", err)
			break
		}

		err = binary.Read(conn, binary.BigEndian, &count)
		if err != nil {
			log.Printf("Buffer: %d\n", err)
			break
		}
		log.Printf("Receive From Client %d\n", count)

		count++
	}

	if err := conn.Close(); err != nil {
		log.Printf("Close: %v\n", err)
	}
}