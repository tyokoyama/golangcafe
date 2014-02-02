package main

import (
	"encoding/binary"
	"log"
	"net"
)

func main() {
	ch := make(chan int)

	for i := 1; i < 4; i++ {
		go connectGoroutine(ch, i)
	}

	count := 0
	for i := range ch {
		log.Printf("Connect End: %d\n", i)
		count++
		if count >= 3 { break }
	}

	close(ch)
}

func connectGoroutine(ch chan<- int, pos int) {
	var recvData uint32

	conn, err := net.Dial("tcp", "localhost:22000")
	if err != nil {
		log.Fatalln(err)
	}

	for i := 0; i < 2; i++ {
		err = binary.Read(conn, binary.LittleEndian, &recvData)
		if err != nil {
			log.Printf("Receive: %s\n", err)
			break
		}
		log.Printf("Receive From Server %v\n", recvData)

		log.Printf("Send To Server %d\n", recvData + 1)
		sendData := uint32(recvData + 1)

		err = binary.Write(conn, binary.BigEndian, sendData)
		if err != nil {
			log.Printf("Send: %v\n", err)
			break
		}
	}

	if err = conn.Close(); err != nil {
		log.Printf("Close: %v\n", err)
	}

	ch <- pos
}
