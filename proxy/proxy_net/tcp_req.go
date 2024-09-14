package proxy_net

import "bronya.com/proxy/iproxy_net"

type TcpReq struct {
	tcpConn iproxy_net.ITcpConn
	packet  []byte
}

func (tcpReq *TcpReq) GetConn() iproxy_net.ITcpConn {
	return tcpReq.tcpConn
}

func (tcpReq *TcpReq) GetPacket() []byte {
	return tcpReq.packet
}
