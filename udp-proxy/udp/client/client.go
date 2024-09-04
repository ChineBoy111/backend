package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	udpConn, err := net.DialUDP(
		"udp", nil, &net.UDPAddr{
			IP:   net.IPv4(127, 0, 0, 1),
			Port: 3302,
		},
	)
	if err != nil {
		os.Exit(1)
	}
	defer func(udpConn *net.UDPConn) {
		if err := udpConn.Close(); err != nil {
			fmt.Println(err)
		}
	}(udpConn)
	for i := 0; i < 3; i++ {
		_, err = udpConn.Write([]byte("client ==> UDP request ==> remote"))
		if err != nil {
			fmt.Println(err)
		}
		var recvBuf [64]byte
		nBytes, remoteAddr, err := udpConn.ReadFromUDP(recvBuf[:])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf(
			"\nNanosecond timestamp: %v\n"+
				"Recieved message: %v\n"+
				"Bytes count: %v\n"+
				"Server address: %v\n",
			time.Now().UnixNano(), string(recvBuf[:nBytes]), nBytes, remoteAddr,
		)
	}
}
