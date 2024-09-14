package iproxy_net

type ITcpMiddleware interface {

	// PreHandler PacketHandler 前的 hook 方法
	PreHandler(tcpRequest ITcpRequest)

	// PacketHandler 处理收到的 tcp 数据包
	PacketHandler(tcpRequest ITcpRequest)

	// PostHandler PacketHandler 后的 hook 方法
	PostHandler(tcpRequest ITcpRequest)
}
