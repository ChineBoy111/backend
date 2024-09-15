package iproxy_net

type ITcpReq interface {
	// GetConn 获取 tcpConn 对象
	GetConn() ITcpConn

	// GetPacket 获取收到的 tcp 数据包
	GetPacket() []byte
}
