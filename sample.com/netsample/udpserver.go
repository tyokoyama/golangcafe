package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
)

func main() {

	udpAddr, err := net.ResolveUDPAddr("udp", ":22000")
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

func receiveGoroutine(conn *net.UDPConn, ch chan<- int) {
	var count uint32 = 1

	for i := 0; i < 2; i++ {
		var data [1024]byte

		_, addr, err := conn.ReadFrom(data[:])
		if err != nil {
			log.Printf("Recv: %v", err)
		}

		buf := bytes.NewBuffer(data[:])
		err = binary.Read(buf, binary.BigEndian, &count)
		if err != nil {
			log.Printf("Buffer: %v\n", err)
			break
		}
		log.Printf("Receive From Client %d\n", count)

		count++

		log.Printf("Send To Client %d\n", count)
		bufW := new(bytes.Buffer)
		err = binary.Write(bufW, binary.LittleEndian, count)
		if err != nil {
			log.Printf("Buffer: %v\n", err)
		}
		_, err = conn.WriteTo(bufW.Bytes(), addr)
		if err != nil {
			log.Printf("Send: %v\n", err)
			break
		}

	}

	if err := conn.Close(); err != nil {
		log.Printf("Close: %v\n", err)
	}

	ch <- 1
}