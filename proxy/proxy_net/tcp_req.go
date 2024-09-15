package proxy_net

import "bronya.com/proxy/iproxy_net"

type TcpReq struct {
	Conn   iproxy_net.ITcpConn
	Packet []byte
}

func (req *TcpReq) GetTcpConn() iproxy_net.ITcpConn {
	return req.Conn
}

func (req *TcpReq) GetPacket() []byte {
	return req.Packet
}
