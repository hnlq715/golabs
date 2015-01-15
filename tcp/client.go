package main

import (
	"bytes"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	conn, _ := net.Dial("tcp", "172.19.108.101:9527")
	tc := conn.(*net.TCPConn)
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(1 * time.Second)
	tc.SetNoDelay(true)
	log.SetFlags(log.LstdFlags)
	n, err := conn.Write([]byte("Hello world!"))
	if err == nil {
		log.Println("Write", n)
	} else {
		log.Println("Write failed", err)
	}

	time.Sleep(4 * time.Second)

	read := bytes.NewReader([]byte("Hello world!"))
	_, err = io.Copy(conn, read)
	if err == nil {
		log.Println("Write", n)
	} else {
		log.Println("Write failed", err)
	}

	//conn.Close()
	buffer := make([]byte, 1)
	n, err = conn.Read(buffer[0:1])
	if err == nil {
		log.Println("Read", n)
	} else {
		log.Println("Read failed", err)
		conn.Close()
	}

	time.Sleep(50 * time.Second)

}
