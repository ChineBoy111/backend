package main

import (
	"bronya.com/proxy/iproxy_net"
	"bronya.com/proxy/proxy_net"
	"log"
)

type PingMiddleware struct {
	proxy_net.TcpBaseMiddleware
}

func (middleware *PingMiddleware) PacketHandler(request iproxy_net.ITcpRequest) {
	_ /* writeBytes */, err := request.GetConnector().GetConn().Write([]byte("ping"))
	if err != nil {
		log.Println(err)
	}
}
func main() {
	server := proxy_net.NewTcpServer("WAN proxy")
	server.SetMiddleware(&PingMiddleware{})
	server.Serve()
}
