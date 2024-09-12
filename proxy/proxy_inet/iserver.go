package proxy_inet

// IServer 服务器接口
type IServer interface {
	Start() // 启动服务器
	Stop()  // 停止服务器
	Serve() // 运行 TCP 服务
}
