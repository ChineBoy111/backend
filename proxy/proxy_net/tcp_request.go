package proxy_net

import "bronya.com/proxy/iproxy_net"

type TcpRequest struct {
	connector iproxy_net.ITcpConnector
	packet    []byte
}

func (tcpRequest *TcpRequest) GetConnector() iproxy_net.ITcpConnector {
	return tcpRequest.connector
}

func (tcpRequest *TcpRequest) GetPacket() []byte {
	return tcpRequest.packet
}
