package inetwork

type TcpMidWareMan interface {
	// UseMidWare 使用中间件
	UseMidWare(msgId uint32, midWare ITcpMidWare)

	// AddMidWare 添加中间件
	AddMidWare(msgId uint32, midWare ITcpMidWare)
}
