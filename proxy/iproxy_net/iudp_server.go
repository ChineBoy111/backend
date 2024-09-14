package iproxy_net

// IUdpServer udp 服务器接口
type IUdpServer interface {
	// Start 启动 udp 服务器
	Start()
	// Serve 运行 udp 服务
	Serve()
	// Stop 停止 udp 服务器
	Stop()
}
