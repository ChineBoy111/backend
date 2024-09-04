package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func serve(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			os.Exit(1)
		}
	}(conn)
	var recvBuf [64]byte
	reader := bufio.NewReader(conn)
	nBytes, err := reader.Read(recvBuf[:])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(
		"\nNanosecond timestamp: %v\n"+
			"Recieved message: %v\n"+
			"Bytes count: %v\n",
		time.Now().UnixNano(), string(recvBuf[:nBytes]), nBytes)
	_, err = conn.Write([]byte("server ==> TCP response ==> remote"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:3302")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("TCP server listening on port 3302")
	for {
		conn, err := listener.Accept() // setup tcp connection
		if err != nil {
			fmt.Println(err)
		}
		go serve(conn) // start a goroutine for handling tcp connection
	}
}
