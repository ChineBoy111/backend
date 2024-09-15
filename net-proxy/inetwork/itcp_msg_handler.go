package inetwork

type ITcpMsgHandler interface {
	// UseMidWare 使用中间件
	UseMidWare(msgId uint32, midWare ITcpMidWare)

	// DoMidWare msg 运行中间件
	DoMidWare(req ITcpReq)
}
