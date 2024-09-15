package main

import (
	"bronya.com/proxy/iproxy_net"
	"bronya.com/proxy/proxy_net"
	"log"
)

type PingMiddleware struct {
	proxy_net.TcpBaseMidware
}

func (middleware *PingMiddleware) Handler(request iproxy_net.ITcpReq) {
	_ /* writeBytes */, err := request.GetTcpConn().GetSocket().Write([]byte("ping"))
	if err != nil {
		log.Println(err)
	}
}

func main() {
	server := proxy_net.NewTcpServer()
	server.SetMidware(&PingMiddleware{})
	server.Serve()
}
