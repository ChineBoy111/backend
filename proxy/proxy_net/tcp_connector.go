package proxy_net

import (
	"bronya.com/proxy/iproxy_net"
	"log"
	"net"
)

// TcpConnector 实现 ITcpConnector 接口
type TcpConnector struct {
	Conn          *net.TCPConn             // tcp 连接对象
	ConnId        uint32                   // tcp 连接 id
	NotifyClosed  chan struct{}            // 通知 tcp 连接已关闭的通道
	packetHandler iproxy_net.PacketHandler // 处理收到的 tcp 数据包的方法
	isClosed      bool                     // tcp 连接是否已关闭
}

// NewTcpConnector 创建 TcpConnector 实例
func NewTcpConnector(conn *net.TCPConn, connId uint32, handler iproxy_net.PacketHandler) *TcpConnector {
	tcpConnector := &TcpConnector{
		Conn:          conn,
		ConnId:        connId,
		NotifyClosed:  make(chan struct{}, 1),
		packetHandler: handler,
		isClosed:      false,
	}
	return tcpConnector
}

// Start 启动 tcp 连接
func (tcpConnector *TcpConnector) Start() {
	log.Println("Start tcpConnector, connId =", tcpConnector.ConnId)
	//! 负责从 tcpConnector.conn 中读的 goroutine
	go tcpConnector.StartReader()
	//! 负责向 tcpConnector.conn 中写的 goroutine
	// go tcpConnector.StartWriter()
}

// StartReader 启动从 tcpConnector.conn 中读的 goroutine
func (tcpConnector *TcpConnector) StartReader() {
	log.Printf("[connId = %v] Start reader goroutine, remoteAddr = %v\n", tcpConnector.ConnId, tcpConnector.GetRemoteAddr())
	defer log.Printf("[connId = %v] Stop reader goroutine, remoteAddr = %v\n", tcpConnector.ConnId, tcpConnector.GetRemoteAddr())
	defer tcpConnector.Stop()

	for {
		// 从 tcpConnector.conn 中读出数据到 buf 中，一次读取 512 字节
		buf := make([]byte, 512)
		readBytes, err := tcpConnector.Conn.Read(buf)
		if err != nil {
			log.Printf("[connId = %v] Read err %v\n", tcpConnector.ConnId, err.Error())
			continue
		}
		// 处理收到的 tcp 数据包
		err = tcpConnector.packetHandler(tcpConnector.Conn, buf, readBytes)
		if err != nil {
			log.Printf("[connId = %v] packetHandler err %v\n", tcpConnector.ConnId, err.Error())
			break
		}
	}
}

// Stop 停止 tcp 连接
func (tcpConnector *TcpConnector) Stop() {
	log.Printf("[connId = %v] Stop tcpConnector\n", tcpConnector.ConnId)
	if tcpConnector.isClosed {
		return
	}
	tcpConnector.isClosed = true
	err := tcpConnector.Conn.Close()
	if err != nil {
		log.Printf("[connId = %v] Stop tcpConnector err\n", tcpConnector.ConnId)
	}
	close(tcpConnector.NotifyClosed)
}

// GetConn 获取 tcp 连接对象
func (tcpConnector *TcpConnector) GetConn() *net.TCPConn {
	return tcpConnector.Conn
}

// GetConnId 获取 tcp 连接 id
func (tcpConnector *TcpConnector) GetConnId() uint32 {
	return tcpConnector.ConnId
}

// GetRemoteAddr 获取客户端的 ip 地址和端口
func (tcpConnector *TcpConnector) GetRemoteAddr() net.Addr {
	return tcpConnector.Conn.RemoteAddr()
}

// Send 发送 tcp 数据包给客户端
func (tcpConnector *TcpConnector) Send() {

}
