package iproxy_net

type ITcpReq interface {
	// GetTcpConn 获取 tcpConn
	GetTcpConn() ITcpConn

	// GetPacket 获取收到的 tcp 数据包
	GetPacket() []byte
}
