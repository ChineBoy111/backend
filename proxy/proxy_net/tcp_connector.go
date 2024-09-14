package proxy_net

import (
	"bronya.com/proxy/iproxy_net"
	"log"
	"net"
)

// TcpConnector 实现 ITcpConnector 接口
type TcpConnector struct {
	Conn         *net.TCPConn              // tcp 套接字
	ConnId       uint32                    // tcp 连接 id
	NotifyClosed chan struct{}             // 通知 tcp 连接已关闭的通道
	isClosed     bool                      // tcp 连接是否已关闭
	Middleware   iproxy_net.ITcpMiddleware // tcp 服务中间件
}

// NewTcpConnector 创建 TcpConnector 实例
func NewTcpConnector(conn *net.TCPConn, connId uint32, middleware iproxy_net.ITcpMiddleware) *TcpConnector {
	tcpConnector := &TcpConnector{
		Conn:         conn,
		ConnId:       connId,
		NotifyClosed: make(chan struct{}, 1),
		isClosed:     false,
		Middleware:   middleware,
	}
	return tcpConnector
}

// Start 启动 tcp 连接
func (connector *TcpConnector) Start() {
	log.Println("Start connector, connId =", connector.ConnId)
	//! 负责从 connector.conn 中读的 goroutine
	go connector.StartReader()
	//! 负责向 connector.conn 中写的 goroutine
	// go connector.StartWriter()
}

// StartReader 启动从 connector.conn 中读的 goroutine
func (connector *TcpConnector) StartReader() {
	log.Printf("[connId = %v] Start reader goroutine, remoteAddr = %v\n", connector.ConnId, connector.GetRemoteAddr())
	defer log.Printf("[connId = %v] Stop reader goroutine, remoteAddr = %v\n", connector.ConnId, connector.GetRemoteAddr())
	defer connector.Stop()

	for {
		// 从 connector.conn 中读出数据到 buf 中，一次读取 512 字节
		buf := make([]byte, 512)
		_ /* readBytes */, err := connector.Conn.Read(buf)
		if err != nil {
			log.Printf("[connId = %v] Read err %v\n", connector.ConnId, err.Error())
			continue
		}

		request := TcpRequest{
			Connector: connector,
			Packet:    buf,
		}

		// 启动使用 tcp 服务中间件的 goroutine，处理收到的 tcp 数据包
		go func(request iproxy_net.ITcpRequest) {
			connector.Middleware.PreHandler(request)
			connector.Middleware.PacketHandler(request)
			connector.Middleware.PostHandler(request)
		}(&request)
	}
}

// Stop 停止 tcp 连接
func (connector *TcpConnector) Stop() {
	log.Printf("[connId = %v] Stop connector\n", connector.ConnId)
	if connector.isClosed {
		return
	}
	connector.isClosed = true
	err := connector.Conn.Close()
	if err != nil {
		log.Printf("[connId = %v] Stop connector err\n", connector.ConnId)
	}
	close(connector.NotifyClosed)
}

// GetConn 获取 tcp 套接字
func (connector *TcpConnector) GetConn() *net.TCPConn {
	return connector.Conn
}

// GetConnId 获取 tcp 连接 id
func (connector *TcpConnector) GetConnId() uint32 {
	return connector.ConnId
}

// GetRemoteAddr 获取客户端的 ip 地址和端口
func (connector *TcpConnector) GetRemoteAddr() net.Addr {
	return connector.Conn.RemoteAddr()
}

// Send 发送 tcp 数据包给客户端
func (connector *TcpConnector) Send() {

}
