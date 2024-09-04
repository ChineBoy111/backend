package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	udpConn, err := net.ListenUDP(
		"udp", &net.UDPAddr{
			IP:   net.IPv4(0, 0, 0, 0),
			Port: 3302,
		},
	)
	if err != nil {
		os.Exit(1)
	}
	defer func(udpConn *net.UDPConn) {
		err := udpConn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(udpConn)
	fmt.Println("UDP server listening on port 3302")
	for {
		var recvBuf [64]byte
		nBytes, remoteAddr, err := udpConn.ReadFromUDP(recvBuf[:])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf(
			"\nNanosecond timestamp: %v\n"+
				"Recieved message: %v\n"+
				"Bytes count: %v\n"+
				"Client address: %v\n",
			time.Now().UnixNano(), string(recvBuf[:nBytes]), nBytes, remoteAddr,
		)
		_, err = udpConn.WriteToUDP([]byte("server ==> UDP response ==> remote"), remoteAddr)
		if err != nil {
			fmt.Println(err)
		}
	}
}
