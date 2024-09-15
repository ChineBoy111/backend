package proxy_net

import "bronya.com/proxy/iproxy_net"

type TcpReq struct {
	Conn   iproxy_net.ITcpConn
	Packet []byte
}

// GetConn 获取 tcpConn 对象
func (req *TcpReq) GetConn() iproxy_net.ITcpConn {
	return req.Conn
}

// GetPacket 获取收到的 tcp 数据包
func (req *TcpReq) GetPacket() []byte {
	return req.Packet
}
