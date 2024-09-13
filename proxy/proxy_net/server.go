package proxy_net

import (
	"bronya.com/proxy/proxy_inet"
	"fmt"
	"log"
	"net"
)

// Server 实现 IServer 接口
type Server struct {
	// 服务器名
	Name string
	// ip 版本
	IpVer string
	// 监听的 ip 地址
	Ip string
	// 监听的端口
	Port int
}

// Start 启动服务器
func (server *Server) Start() {
	log.Printf("Start server %v, ip %v, port %v\n", server.Name, server.Ip, server.Port)
	go func() { //! 监听所有 ip 地址的 tcp 连接请求的 goroutine

		//! 解析 tcp 地址
		tcpAddr, err := net.ResolveTCPAddr(server.IpVer, fmt.Sprintf("%v:%v", server.Ip, server.Port))
		if err != nil {
			log.Println("Resolve tcp addr err", err.Error())
			return
		}

		//! 监听所有 ip 地址的 tcp 连接请求
		tcpListener, err := net.ListenTCP(server.IpVer, tcpAddr)
		if err != nil {
			log.Println("Listen tcp err", err.Error())
			return
		}
		log.Printf("Start server %v ok, listening %v:%v\n", server.Name, server.Ip, server.Port)

		var connId uint32 = 0
		//! 阻塞等待客户端的 tcp 连接请求
		for {
			tcpConn, err := tcpListener.AcceptTCP() // 收到客户端的 tcp 连接请求
			if err != nil {
				log.Println("Accept tcp err", err.Error())
				continue
			}

			conn := NewConn(tcpConn, connId, DataHandler)
			connId++
			go conn.Start() //! 处理 tcp 连接的 goroutine
		}
	}()
}

func DataHandler(tcpConn *net.TCPConn, buf []byte, bytesCnt int) error {
	fmt.Println("========== Echo to client after codec and zlib ==========")
	_ /* writeBytes */, err := tcpConn.Write(buf[:bytesCnt])
	if err != nil {
		log.Println("Write err", err.Error())
	}
	return err
}

// Serve 运行服务
func (server *Server) Serve() {
	//! 启动服务器，运行服务
	server.Start()
	//TODO

	//! ==================== 阻塞等待 ====================
	select {}
}

// Stop 停止服务器
func (server *Server) Stop() {
	//TODO
}

// NewServer 创建 Server 实例
func NewServer(name string) proxy_inet.IServer {
	server := Server{
		Name:  name,      // 服务器名
		IpVer: "tcp4",    // ip 版本
		Ip:    "0.0.0.0", // 监听所有 ip 地址的 tcp 连接请求
		Port:  3333,      // 监听的端口
	}
	return &server
}
