package iproxy_net

import "net"

type ITcpConnector interface {

	// Start 启动 tcp 连接
	Start()

	// Stop 停止 tcp 连接
	Stop()

	// GetConn 获取 tcp 套接字
	GetConn() *net.TCPConn

	// GetConnId 获取 tcp 连接 id
	GetConnId() uint32

	// GetRemoteAddr 获取客户端的 ip 地址和端口
	GetRemoteAddr() net.Addr

	// Send 发送 tcp 数据包给客户端
	Send()
}
