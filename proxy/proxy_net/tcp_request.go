package proxy_net

import "bronya.com/proxy/iproxy_net"

type TcpRequest struct {
	Connector iproxy_net.ITcpConnector
	Packet    []byte
}

func (request *TcpRequest) GetConnector() iproxy_net.ITcpConnector {
	return request.Connector
}

func (request *TcpRequest) GetPacket() []byte {
	return request.Packet
}
