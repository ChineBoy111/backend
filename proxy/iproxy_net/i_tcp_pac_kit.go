package iproxy_net

type ITcpPacKit interface {
	// GetHeadLen 获取 tcp 数据包的 head 长度
	GetHeadLen() uint32

	// Pack tcp 封包，msg 结构体实例序列化为 packet 字节数组
	Pack(msg ITcpMsg) ([]byte, error)

	// Unpack tcp 拆包，packet 字节数组反序列化为 msg 结构体实例
	Unpack(packet []byte) (ITcpMsg, error)
}
