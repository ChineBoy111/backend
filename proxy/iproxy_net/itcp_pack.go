package iproxy_net

type ITcpPack interface {
	// GetHeadSize 获取 tcp 数据包的 head 头部大小
	GetHeadSize() uint32

	// Pack tcp 封包
	Pack(msg ITcpMsg) ([]byte, error)

	// Unpack tcp 拆包
	Unpack(packet []byte) (ITcpMsg, error)
}
