package proxy_net

import (
	"bronya.com/proxy/iproxy_net"
	"log"
	"net"
)

// TcpConn 实现 ITcpConn 接口
type TcpConn struct {
	Closed   chan struct{}              // 通知 tcp 连接已关闭的通道
	Id       uint32                     // tcp 连接 id
	MidWare  iproxy_net.ITcpBaseMidWare // tcp 服务中间件
	Socket   *net.TCPConn               // tcp 套接字
	isClosed bool                       // tcp 连接是否已关闭
}

// NewTcpConn 创建 TcpConn 实例
func NewTcpConn(socket *net.TCPConn, id uint32, midWare iproxy_net.ITcpBaseMidWare) *TcpConn {
	conn := &TcpConn{
		Closed:   make(chan struct{}, 1),
		Id:       id,
		MidWare:  midWare,
		Socket:   socket,
		isClosed: false,
	}
	return conn
}

// Start 启动 tcp 连接
func (conn *TcpConn) Start() {
	log.Printf("[conn %v] Start tcp conn\n", conn.Id)
	//! 负责从 conn.Socket 中读的 goroutine
	go conn.StartReader()
	//! 负责向 conn.Socket 中写的 goroutine
	// go conn.StartWriter()
}

// StartReader 启动从 conn.Socket 中读的 goroutine
func (conn *TcpConn) StartReader() {
	log.Printf("[conn %v] Start reader goroutine, remoteAddr = %v\n", conn.Id, conn.GetRemoteAddr())
	defer log.Printf("[conn %v] Stop reader goroutine, remoteAddr = %v\n", conn.Id, conn.GetRemoteAddr())
	defer conn.Stop()

	for {
		// 从 conn.Socket 中读出 512 字节的数据到 buf 中
		buf := make([]byte, 512)
		_ /* readBytes */, err := conn.Socket.Read(buf)
		if err != nil {
			log.Printf("[conn %v] Read err: %v\n", conn.Id, err.Error())
			continue
		}

		req := TcpReq{
			Conn:   conn,
			Packet: buf,
		}

		// 启动使用 tcp 服务中间件的 goroutine，处理收到的 tcp 数据包
		go func(req_ iproxy_net.ITcpReq) {
			conn.MidWare.PreHandler(req_)
			conn.MidWare.Handler(req_)
			conn.MidWare.PostHandler(req_)
		}(&req)
	}
}

// Stop 停止 tcp 连接
func (conn *TcpConn) Stop() {
	log.Printf("[conn %v] Stop conn\n", conn.Id)
	if conn.isClosed {
		return
	}
	conn.isClosed = true
	err := conn.Socket.Close()
	if err != nil {
		log.Printf("[conn %v] Stop conn err: %v\n", conn.Id, err.Error())
	}
	close(conn.Closed)
}

// GetId 获取 tcp 连接 id
func (conn *TcpConn) GetId() uint32 {
	return conn.Id
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
