package proxy_inet

import "net"

type IConn interface {
	// Start 启动连接
	Start()
	// Stop 停止连接
	Stop()
	// GetTcpConn 获取 tcp 连接对象（tcp 套接字）
	GetTcpConn() *net.TCPConn
	// GetConnId 获取连接 id
	GetConnId() uint32
	// GetRemoteAddr 获取客户端的 ip 地址和端口
	GetRemoteAddr() net.Addr
	// Send 发送数据给客户端
	Send()
}

// DataHandler 处理收到的数据
type DataHandler func(tcpConn *net.TCPConn, buf []byte, bytesCnt int) error
