package main

import (
	"bronya.com/net-proxy/inetwork"
	"bronya.com/net-proxy/network"
	"log"
)

type PingMidWare struct {
	network.TcpBaseMidWare
}

type EchoMidWare struct {
	network.TcpBaseMidWare
}

func (midWare *PingMidWare) MsgHandler(req inetwork.ITcpReq) {
	log.Printf("Msg: len=%v, id=%v, data=%v\n", req.GetMsgLen(), req.GetMsgId(), string(req.GetMsgData()))
	err := req.GetConn().SendPac(1, []byte("ping"))
	if err != nil {
		log.Println("Send pac err", err)
	}
}

func (midWare *EchoMidWare) MsgHandler(req inetwork.ITcpReq) {
	log.Printf("Msg: len=%v, id=%v, data=%v\n", req.GetMsgLen(), req.GetMsgId(), string(req.GetMsgData()))
	err := req.GetConn().SendPac(1, []byte("echo"))
	if err != nil {
		log.Println("Send pac err", err)
	}
}

func main() {
	// 创建 tcp 服务器
	server := network.NewTcpServer()
	// 使用中间件
	server.BindMidWare(0 /* msgId */, &PingMidWare{} /* midWare */)
	server.BindMidWare(1 /* msgId */, &EchoMidWare{} /* midWare */)
	// 启动 tcp 服务器，运行 tcp 服务
	server.Serve()
}
