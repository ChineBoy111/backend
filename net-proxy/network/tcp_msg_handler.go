package network

import (
	"bronya.com/net-proxy/inetwork"
	"log"
)

type TcpMsgHandler struct {
	BindDict map[uint32]inetwork.ITcpMidWare
}

var msgHandler *TcpMsgHandler

func init() {
	msgHandler = &TcpMsgHandler{
		BindDict: make(map[uint32]inetwork.ITcpMidWare),
	}
}

func NewTcpMsgHandler() *TcpMsgHandler {
	return msgHandler
}

func (m *TcpMsgHandler) BindMidWare(msgId uint32, midWare inetwork.ITcpMidWare) {
	if _, ok := m.BindDict[msgId]; ok {
		log.Println("Use midWare err")
		return
	}
	m.BindDict[msgId] = midWare
	log.Println("Use midWare ok")
}

func (m *TcpMsgHandler) UseMidWare(req inetwork.ITcpReq) {
	m.BindDict[]
}
