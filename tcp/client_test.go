package main

import (
	"log"
	"net"
	"testing"
	"time"
)

func Benchmark_1(b *testing.B) {
	b.SetBytes(1024)
	b.StopTimer()
	b.StartTimer()

	conn, _ := net.Dial("tcp", "172.19.108.101:9527")
	tc := conn.(*net.TCPConn)
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(1 * time.Second)
	tc.SetNoDelay(true)
	log.SetFlags(log.LstdFlags)
	_, err := conn.Write([]byte("Hello world!"))
	if err != nil {
		log.Println("Write failed", err)
	}

	for i := 0; i < b.N; i++ {
		buffer := make([]byte, 1)
		_, err = conn.Read(buffer[0:1])
		if err != nil {
			b.Log("close")
			conn.Close()
		}
	}

}
