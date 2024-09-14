package iproxy_net

type ITcpReq interface {
	// GetConn 获取已建立连接的 tcpConn
	GetConn() ITcpConn
	// GetPacket 获取收到的 tcp 数据包
	GetPacket() []byte
}
