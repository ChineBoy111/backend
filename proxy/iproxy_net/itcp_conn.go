package iproxy_net

import "net"

type ITcpConn interface {
	// Start 启动 tcp 连接
	Start()
	// Stop 停止 tcp 连接
	Stop()
	// GetSocket 获取 tcp 连接对象（tcp 套接字）
	GetSocket() *net.TCPConn
	// GetConnId 获取 tcp 连接 id
	GetConnId() uint32
	// GetRemoteAddr 获取客户端的 ip 地址和端口
	GetRemoteAddr() net.Addr
	// Send 发送 tcp 数据包给客户端
	Send()
}

// TcpPacketHandler 处理收到的 tcp 数据包
type TcpPacketHandler func(conn *net.TCPConn, buf []byte, bytesCnt int) error
