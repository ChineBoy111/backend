package main

import (
	"bronya.com/net-proxy/iproxy"
	"bronya.com/net-proxy/proxy"
	"log"
)

type PingMidWare struct {
	proxy.TcpBaseMidWare
}

func (midWare *PingMidWare) MsgHandler(req iproxy.ITcpReq) {
	log.Printf("Msg: len=%v, id=%v, data=%v\n", req.GetMsgLen(), req.GetMsgId(), string(req.GetMsgData()))
	err := req.GetConn().SendPac(1, []byte("ping"))
	if err != nil {
		log.Println("Send pac err", err)
	}
}

func main() {
	// 创建 tcp 服务器
	server := proxy.NewTcpServer()
	// 使用中间件
	server.SetMidWare(&PingMidWare{})
	// 启动 tcp 服务器，运行 tcp 服务
	server.Serve()
}
