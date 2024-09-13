package proxy_net

import (
	"bronya.com/proxy/proxy_inet"
	"log"
	"net"
)

type Conn struct {
	TCPConn      *net.TCPConn           // TCP 连接对象
	ConnID       uint32                 // 连接 ID
	byteHandler  proxy_inet.ByteHandler // 处理收到的数据
	isClosed     bool                   // 连接是否已关闭
	NotifyClosed chan struct{}          // 通知连接已关闭的通道
}

func NewConn(tcpConn *net.TCPConn, connID uint32, callback proxy_inet.ByteHandler) *Conn {
	conn := &Conn{
		TCPConn:      tcpConn,
		ConnID:       connID,
		byteHandler:  callback,
		isClosed:     false,
		NotifyClosed: make(chan struct{}, 1),
	}
	return conn
}

// Start 启动连接
func (conn *Conn) Start() {
	log.Println("Start conn, connID =", conn.ConnID)
	// 从套接字中读
	go conn.ReaderDo()
	// 向套接字中写
}

func (conn *Conn) ReaderDo() {
	log.Println("Start reader goroutine, connID =", conn.ConnID)
	defer log.Printf("Stop reader goroutine, connID = %v, remoteAddr = %v\n", conn.ConnID, conn.GetRemoteAddr())
	defer conn.Stop()

	for {
		// 从套接字中读出数据到 buf 中，一次读取 512 字节
		buf := make([]byte, 512)
		readBytes, err := conn.TCPConn.Read(buf)
		if err != nil {
			log.Println("Read err", err)
			continue
		}
		// 调用
	}
}

// Stop 停止连接
func (conn *Conn) Stop() {
	log.Println("Stop conn, connID =", conn.ConnID)
	if conn.isClosed {
		return
	}
	conn.isClosed = true
	err := conn.TCPConn.Close()
	if err != nil {
		log.Println("Stop conn err, connID =", conn.ConnID)
	}
	close(conn.NotifyClosed)
}

// GetTCPConn 获取 TCP 连接对象（TCP 套接字）
func (conn *Conn) GetTCPConn() *net.TCPConn {
	return conn.TCPConn
}

// GetConnID 获取连接 ID
func (conn *Conn) GetConnID() uint32 {
	return conn.ConnID
}

// GetRemoteAddr 获取客户端的 IP 地址和端口
func (conn *Conn) GetRemoteAddr() net.Addr {
	return conn.TCPConn.RemoteAddr()
}

// Send 发送数据给客户端
func (conn *Conn) Send() {

}
