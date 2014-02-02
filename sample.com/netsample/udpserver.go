package main

import (
	"encoding/binary"
	"log"
	"net"
)

func main() {

	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:22000")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		ch := make(chan int)

		// ListenUDPはudp/udp4/udp6でなければエラーになる。
		udpConn, err := net.ListenUDP("udp", udpAddr)
		if err != nil {
			log.Printf("ListenUDP: %v\n", err)
			continue
		}

		go receiveGoroutine(udpConn, ch)

		<-ch
	}
}

func receiveGoroutine(conn net.Conn, ch chan<- int) {
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

	ch <- 1
}