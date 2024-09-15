package proxy_net

import (
	"bronya.com/proxy/iproxy_net"
	"log"
	"net"
)

// TcpConn 实现 ITcpConn 接口
type TcpConn struct {
	Closed   chan struct{}              // 通知 tcp 连接已关闭的通道
	ConnId   uint32                     // tcp 连接 id
	Midware  iproxy_net.ITcpBaseMidware // tcp 服务中间件
	Socket   *net.TCPConn               // tcp 套接字
	isClosed bool                       // tcp 连接是否已关闭
}

// NewTcpConn 创建 TcpConn 实例
func NewTcpConn(socket *net.TCPConn, connId uint32, midware iproxy_net.ITcpBaseMidware) *TcpConn {
	tcpConn := &TcpConn{
		Closed:   make(chan struct{}, 1),
		ConnId:   connId,
		Midware:  midware,
		Socket:   socket,
		isClosed: false,
	}
	return tcpConn
}

// Start 启动 tcp 连接
func (conn *TcpConn) Start() {
	log.Printf("[connId = %v] Start tcp conn\n", conn.ConnId)
	//! 负责从 conn.Socket 中读的 goroutine
	go conn.StartReader()
	//! 负责向 conn.Socket 中写的 goroutine
	// go conn.StartWriter()
}

// StartReader 启动从 conn.Socket 中读的 goroutine
func (conn *TcpConn) StartReader() {
	log.Printf("[connId = %v] Start reader goroutine, remoteAddr = %v\n", conn.ConnId, conn.GetRemoteAddr())
	defer log.Printf("[connId = %v] Stop reader goroutine, remoteAddr = %v\n", conn.ConnId, conn.GetRemoteAddr())
	defer conn.Stop()

	for {
		// 从 conn.Socket 中读出 512 字节的数据到 buf 中
		buf := make([]byte, 512)
		_ /* readBytes */, err := conn.Socket.Read(buf)
		if err != nil {
			log.Printf("[connId = %v] Read err: %v\n", conn.ConnId, err.Error())
			continue
		}

		req := TcpReq{
			Conn:   conn,
			Packet: buf,
		}

		// 启动使用 tcp 服务中间件的 goroutine，处理收到的 tcp 数据包
		go func(req_ iproxy_net.ITcpReq) {
			conn.Midware.PreHandler(req_)
			conn.Midware.Handler(req_)
			conn.Midware.PostHandler(req_)
		}(&req)
	}
}

// Stop 停止 tcp 连接
func (conn *TcpConn) Stop() {
	log.Printf("[connId = %v] Stop conn\n", conn.ConnId)
	if conn.isClosed {
		return
	}
	conn.isClosed = true
	err := conn.Socket.Close()
	if err != nil {
		log.Printf("[connId = %v] Stop conn err: %v\n", conn.ConnId, err.Error())
	}
	close(conn.Closed)
}

// GetConnId 获取 tcp 连接 id
func (conn *TcpConn) GetConnId() uint32 {
	return conn.ConnId
}

// GetRemoteAddr 获取客户端的 ip 地址和端口
func (conn *TcpConn) GetRemoteAddr() net.Addr {
	return conn.Socket.RemoteAddr()
}

// GetSocket 获取 tcp 套接字
func (conn *TcpConn) GetSocket() *net.TCPConn {
	return conn.Socket
}

// Send 发送 tcp 数据包给客户端
func (conn *TcpConn) Send() {

}
