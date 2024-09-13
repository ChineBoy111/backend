package proxy_inet

import "net"

type IConn interface {
	// Start 启动连接
	Start()
	// Stop 停止连接
	Stop()
	// GetTCPConn 获取 TCP 连接对象（TCP 套接字）
	GetTCPConn() *net.TCPConn
	// GetConnID 获取连接 ID
	GetConnID() uint32
	// GetRemoteAddr 获取客户端的 IP 地址和端口
	GetRemoteAddr() net.Addr
	// Send 发送数据给客户端
	Send()
}

// ByteHandler 处理收到的数据
type ByteHandler func(*net.TCPConn, []byte, int) error
