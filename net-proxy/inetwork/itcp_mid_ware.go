package inetwork

type ITcpMidWare interface {
	// PreHandler Handler 前的 hook 方法
	PreHandler(req ITcpReq)

	// MsgHandler 处理拆包得到的 tcp 消息
	MsgHandler(req ITcpReq)

	// PostHandler Handler 后的 hook 方法
	PostHandler(req ITcpReq)
}
