package main

import (
	"bronya.com/proxy/iproxy_net"
	"bronya.com/proxy/proxy_net"
	"log"
)

type PingMiddleware struct {
	proxy_net.TcpBaseMidWare
}

func (middleware *PingMiddleware) Handler(request iproxy_net.ITcpReq) {
	_ /* writeBytes */, err := request.GetConn().GetSocket().Write([]byte("ping"))
	if err != nil {
		log.Println("Write err:", err.Error())
	}
}

func main() {
	server := proxy_net.NewTcpServer()
	server.SetMidWare(&PingMiddleware{})
	server.Serve()
}
