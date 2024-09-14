package iproxy_net

// IServer tcp/udp 服务器接口
type IServer interface {
	// Start 启动 tcp/udp 服务器
	Start()
	// Serve 运行 tcp/udp 服务
	Serve()
	// Stop 停止 tcp/udp 服务器
	Stop()
}
