package iproxy_net

type ITcpRequest interface {

	// GetConnector 获取 tcpConnector
	GetConnector() ITcpConnector

	// GetPacket 获取收到的 tcp 数据包
	GetPacket() []byte
}
