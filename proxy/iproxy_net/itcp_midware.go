package iproxy_net

type ITcpBaseMidware interface {
	// PreHandler Handler 前的 hook 方法
	PreHandler(req ITcpReq)

	// Handler 处理收到的 tcp 数据包
	Handler(req ITcpReq)

	// PostHandler Handler 后的 hook 方法
	PostHandler(req ITcpReq)
}
