package iproxy

type ITcpPacKit interface {
	// GetHeadLen 获取 tcp 数据包的 head 长度
	GetHeadLen() uint32

	// Pack 封包，将 msg 结构体变量序列化为 packet 字节数组（tcp 消息 -> tcp 数据包）
	Pack(msg ITcpMsg) ([]byte, error)

	// Unpack 拆包，将 packet 字节数组反序列化为 msg 结构体变量（tcp 数据包 -> tcp 消息）
	Unpack(pac []byte) (ITcpMsg, error)
}
