package main

import (
	"log"
	"net"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":9527")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	log.Println("Listening...")

	for {
		buffer := make([]byte, 1024)
		log.Println("Accepting...")
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			log.Println("tcp accept failed!")
			break
		} else {
			n, err := tcpConn.Read(buffer[0:])
			if err == nil {
				log.Println("Read", n, buffer)
			} else {
				log.Println("Read error", err)
			}
		}
		//tcpConn.Close()
	}
}
