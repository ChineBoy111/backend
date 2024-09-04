package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3302")
	if err != nil {
		os.Exit(1)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)
	for i := 0; i < 3; i++ {
		_, err = conn.Write([]byte("client ==> TCP request ==> remote"))
		if err != nil {
			fmt.Println(err)
		}
		var recvBuf [64]byte
		nBytes, err := conn.Read(recvBuf[:])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf(
			"\nNanosecond timestamp: %v\n"+
				"Recieved message: %v\n"+
				"Bytes count: %v\n",
			time.Now().UnixNano(), string(recvBuf[:nBytes]), nBytes,
		)
	}
}
