package proxy_net

import (
	"bronya.com/proxy/iproxy_net"
	"fmt"
	"log"
	"net"
)

// TcpServer 实现 ITcpServer 接口
type TcpServer struct {
	Name       string                    // tcp 服务器名
	IpVer      string                    // ip 版本
	Ip         string                    // 监听的 ip 地址
	Port       int                       // 监听的端口
	Middleware iproxy_net.ITcpMiddleware // tcp 服务中间件
}

// Start 启动 tcp 服务器
func (server *TcpServer) Start() {
	log.Printf("Start server %v, ip %v, port %v\n", server.Name, server.Ip, server.Port)
	go func() { //! 负责监听所有 ip 地址的 tcp 连接请求的 goroutine

		//! 解析 tcp 地址
		tcpAddr, err := net.ResolveTCPAddr(server.IpVer, fmt.Sprintf("%v:%v", server.Ip, server.Port))
		if err != nil {
			log.Println("ResolveTCPAddr err", err.Error())
			return
		}

		//! 监听所有 ip 地址的 tcp 连接请求
		tcpListener, err := net.ListenTCP(server.IpVer, tcpAddr)
		if err != nil {
			log.Println("ListenTCP err", err.Error())
			return
		}
		log.Printf("Start server %v ok, listening %v:%v\n", server.Name, server.Ip, server.Port)

		var connId uint32 = 0
		//! 阻塞等待客户端的 tcp 连接请求
		for {
			conn, err := tcpListener.AcceptTCP() // 收到客户端的 tcp 连接请求
			if err != nil {
				log.Println("AcceptTCP err", err.Error())
				continue
			}

			//! 已建立连接的 tcpConn
			tcpConn := NewTcpConnector(conn, connId, server.Middleware)
			connId++

			go tcpConn.Start() //! 负责处理 tcp 连接的 goroutine
		}
	}()
}

// Serve 运行 tcp 服务
func (server *TcpServer) Serve() {
	//! 启动 tcp 服务器，运行服务
	server.Start()
	//TODO

	//! ==================== 阻塞等待 ====================
	select {}
}

// Stop 停止 tcp 服务器
func (server *TcpServer) Stop() {
	//TODO
}

// SetMiddleware 设置 tcp 服务中间件
func (server *TcpServer) SetMiddleware(middleware iproxy_net.ITcpMiddleware) {
	server.Middleware = middleware
}

// NewTcpServer 创建 TcpServer 实例
func NewTcpServer(name string) iproxy_net.ITcpServer {
	server := TcpServer{
		Name:       name,      // 服务器名
		IpVer:      "tcp4",    // ip 版本
		Ip:         "0.0.0.0", // 监听所有 ip 地址的 tcp 连接请求
		Port:       3333,      // 监听的端口
		Middleware: nil,
	}
	return &server
}
