package proxy_inet

// IServer 服务器接口
type IServer interface {
	// Start 启动服务
	Start()
	// Serve 运行服务
	Serve()
	// Stop 停止服务
	Stop()
}
