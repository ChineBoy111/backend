package main

import (
	"bronya.com/proxy/iproxy_net"
	"bronya.com/proxy/proxy_net"
	"log"
)

type PingMidWare struct {
	proxy_net.TcpBaseMidWare
}

func (midWare *PingMidWare) MsgHandler(request iproxy_net.ITcpReq) {
	_ /* writeBytes */, err := request.GetConn().GetSocket().Write([]byte("ping"))
	if err != nil {
		log.Println("Write err", err.Error())
	}
}

func main() {
	server := proxy_net.NewTcpServer()
	server.SetMidWare(&PingMidWare{})
	server.Serve()
}
