package iproxy_net

// ITcpServer tcp 服务器接口
type ITcpServer interface {
	// Start 启动 tcp 服务器
	Start()

	// Serve 运行 tcp 服务
	Serve()

	// Stop 停止 tcp 服务器
	Stop()

	// SetMidWare 设置 tcp 服务中间件
	SetMidWare(midWare ITcpBaseMidWare)
}
