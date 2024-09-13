package proxy_net

import (
	"bronya.com/proxy/proxy_inet"
	"log"
	"net"
)

type Conn struct {
	TcpConn      *net.TCPConn           // tcp 连接对象
	ConnId       uint32                 // 连接 id
	NotifyClosed chan struct{}          // 通知连接已关闭的通道
	dataHandler  proxy_inet.DataHandler // 处理收到的数据
	isClosed     bool                   // 连接是否已关闭
}

func NewConn(tcpConn *net.TCPConn, connId uint32, dataHandler_ proxy_inet.DataHandler) *Conn {
	conn := &Conn{
		//* public
		TcpConn:      tcpConn,
		ConnId:       connId,
		NotifyClosed: make(chan struct{}, 1),
		//* private
		dataHandler: dataHandler_,
		isClosed:    false,
	}
	return conn
}

// Start 启动连接
func (conn *Conn) Start() {
	log.Println("Start conn, connId =", conn.ConnId)
	//! 从 conn.tcpConn 套接字中读的 goroutine
	go conn.StartReader()
	//! 向 conn.tcpConn 套接字中写的 goroutine
	// go conn.StartWriter()
}

// StartReader 启动从 conn.tcpConn 套接字中读的 goroutine
func (conn *Conn) StartReader() {
	log.Printf("[connId = %v] Start reader goroutine, remoteAddr = %v\n", conn.ConnId, conn.GetRemoteAddr())
	defer log.Printf("[connId = %v] Stop reader goroutine, remoteAddr = %v\n", conn.ConnId, conn.GetRemoteAddr())
	defer conn.Stop()

	for {
		// 从 tcp 套接字中读出数据到 buf 中，一次读取 512 字节
		buf := make([]byte, 512)
		readBytes, err := conn.TcpConn.Read(buf)
		if err != nil {
			log.Printf("[connId = %v] Read err %v\n", conn.ConnId, err.Error())
			continue
		}
		// 处理收到的数据
		err = conn.dataHandler(conn.TcpConn, buf, readBytes)
		if err != nil {
			log.Printf("[connId = %v] DataHandler err %v\n", conn.ConnId, err.Error())
			break
		}
	}
}

// Stop 停止连接
func (conn *Conn) Stop() {
	log.Printf("[connId = %v] Stop conn\n", conn.ConnId)
	if conn.isClosed {
		return
	}
	conn.isClosed = true
	err := conn.TcpConn.Close()
	if err != nil {
		log.Printf("[connId = %v] Stop conn err\n", conn.ConnId)
	}
	close(conn.NotifyClosed)
}

// GetTcpConn 获取 tcp 连接对象（tcp 套接字）
func (conn *Conn) GetTcpConn() *net.TCPConn {
	return conn.TcpConn
}

// GetConnId 获取连接 id
func (conn *Conn) GetConnId() uint32 {
	return conn.ConnId
}

// GetRemoteAddr 获取客户端的 ip 地址和端口
func (conn *Conn) GetRemoteAddr() net.Addr {
	return conn.TcpConn.RemoteAddr()
}

// Send 发送数据给客户端
func (conn *Conn) Send() {
}
