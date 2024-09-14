package iproxy_net

type ITcpMiddleware interface {
	// PreHandler TcpConnector.packetHandler 前的 hook 方法
	PreHandler(tcpRequest ITcpRequest)

	Handler(tcpRequest ITcpRequest)

	// PostHandler TcpConnector.packetHandler 后的 hook 方法
	PostHandler(tcpRequest ITcpRequest)
}
