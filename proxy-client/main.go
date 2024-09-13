package main

import (
	"log"
	"net"
	"time"
)

func main() {
	log.Println("Start client")
	time.Sleep(1 * time.Second)

	//! 连接到指定的 IP 地址
	conn, err := net.Dial("tcp4", "127.0.0.1:3333")
	if err != nil {
		log.Fatalf("Start client error %s\n", err.Error())
	}

	for {
		_ /* writeBytes */, err := conn.Write([]byte("Hello wanproxy!"))
		if err != nil {
			log.Fatalf("Write error %s\n", err.Error())
		}

		buf := make([]byte, 64)
		readBytes, err := conn.Read(buf)
		if err != nil {
			log.Fatalf("Read error %s\n", err.Error())
		}
		log.Printf("Read %d bytes\n", readBytes)
		time.Sleep(1 * time.Second)
	}
}
