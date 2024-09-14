package proxy_net

import (
	"bronya.com/proxy/iproxy_net"
	"log"
	"net"
)

// TcpConn 实现 ITcpConn 接口
type TcpConn struct {
	Socket           *net.TCPConn                // tcp 连接对象（tcp 套接字）
	ConnId           uint32                      // tcp 连接 id
	NotifyClosed     chan struct{}               // 通知 tcp 连接已关闭的通道
	tcpPacketHandler iproxy_net.TcpPacketHandler // 处理收到的 tcp 数据包的方法
	isClosed         bool                        // tcp 连接是否已关闭
}

// NewTcpConn 创建 TcpConn 实例
func NewTcpConn(socket *net.TCPConn, connId uint32, callback iproxy_net.TcpPacketHandler) *TcpConn {
	tcpConn := &TcpConn{
		Socket:           socket,
		ConnId:           connId,
		NotifyClosed:     make(chan struct{}, 1),
		tcpPacketHandler: callback,
		isClosed:         false,
	}
	return tcpConn
}

// Start 启动 tcp 连接
func (tcpConn *TcpConn) Start() {
	log.Println("Start tcpConn, connId =", tcpConn.ConnId)
	//! 负责从 tcpConn.socket 套接字中读的 goroutine
	go tcpConn.StartReader()
	//! 负责向 tcpConn.socket 套接字中写的 goroutine
	// go tcpConn.StartWriter()
}

// StartReader 启动从 tcpConn.socket 套接字中读的 goroutine
func (tcpConn *TcpConn) StartReader() {
	log.Printf("[connId = %v] Start reader goroutine, remoteAddr = %v\n", tcpConn.ConnId, tcpConn.GetRemoteAddr())
	defer log.Printf("[connId = %v] Stop reader goroutine, remoteAddr = %v\n", tcpConn.ConnId, tcpConn.GetRemoteAddr())
	defer tcpConn.Stop()

	for {
		// 从 tcpConn.socket 套接字中读出数据到 buf 中，一次读取 512 字节
		buf := make([]byte, 512)
		readBytes, err := tcpConn.Socket.Read(buf)
		if err != nil {
			log.Printf("[connId = %v] Read err %v\n", tcpConn.ConnId, err.Error())
			continue
		}
		// 处理收到的 tcp 数据包
		err = tcpConn.tcpPacketHandler(tcpConn.Socket, buf, readBytes)
		if err != nil {
			log.Printf("[connId = %v] tcpPacketHandler err %v\n", tcpConn.ConnId, err.Error())
			break
		}
	}
}

// Stop 停止 tcp 连接
func (tcpConn *TcpConn) Stop() {
	log.Printf("[connId = %v] Stop tcpConn\n", tcpConn.ConnId)
	if tcpConn.isClosed {
		return
	}
	tcpConn.isClosed = true
	err := tcpConn.Socket.Close()
	if err != nil {
		log.Printf("[connId = %v] Stop tcpConn err\n", tcpConn.ConnId)
	}
	close(tcpConn.NotifyClosed)
}

// GetSocket 获取 tcp 连接对象（tcp 套接字）
func (tcpConn *TcpConn) GetSocket() *net.TCPConn {
	return tcpConn.Socket
}

// GetConnId 获取 tcp 连接 id
func (tcpConn *TcpConn) GetConnId() uint32 {
	return tcpConn.ConnId
}

// GetRemoteAddr 获取客户端的 ip 地址和端口
func (tcpConn *TcpConn) GetRemoteAddr() net.Addr {
	return tcpConn.Socket.RemoteAddr()
}

// Send 发送 tcp 数据包给客户端
func (tcpConn *TcpConn) Send() {

}
