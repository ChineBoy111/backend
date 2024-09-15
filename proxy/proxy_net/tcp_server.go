package proxy_net

import (
	"bronya.com/proxy/iproxy_net"
	"bronya.com/proxy/utils"
	"fmt"
	"log"
	"net"
)

// TcpServer 实现 ITcpServer 接口
type TcpServer struct {
	Proto   string                     // 协议
	HostIp  string                     // 监听 HostIp 的 tcp 连接请求
	Port    int                        // 监听的端口
	MidWare iproxy_net.ITcpBaseMidWare // tcp 服务中间件
}

// Start 启动 tcp 服务器
func (server *TcpServer) Start() {
	log.Printf("Start server %v, ver %v\n", utils.Global.Name, utils.Global.Ver)
	log.Println("Copyright (c) bronya.com")
	log.Println("All rights reserved")
	go func() { //! 负责监听 HostIp 的 tcp 连接请求的 goroutine

		//! 解析 tcp 地址
		tcpAddr, err := net.ResolveTCPAddr(server.Proto, fmt.Sprintf("%v:%v", server.HostIp, server.Port))
		if err != nil {
			log.Println("Resolve tcp addr err", err.Error())
			return
		}

		//! 监听 HostIp 的 tcp 连接请求
		tcpListener, err := net.ListenTCP(server.Proto, tcpAddr)
		if err != nil {
			log.Println("Listen tcp err", err.Error())
			return
		}
		log.Printf("Server listening %v:%v\n", server.HostIp, server.Port)

		var id uint32 = 0
		//! 阻塞等待客户端的 tcp 连接请求
		for {
			conn, err := tcpListener.AcceptTCP() // 收到客户端的 tcp 连接请求
			if err != nil {
				log.Println("Accept tcp err", err.Error())
				continue
			}

			//! 已建立连接的 tcpConn
			tcpConn := NewTcpConn(conn, id, server.MidWare)
			id++

			go tcpConn.Start() //! 负责处理 tcp 连接的 goroutine
		}
	}()
}

// Serve 运行 tcp 服务
func (server *TcpServer) Serve() {
	//! 启动 tcp 服务器，运行服务
	server.Start()
	select {} //todo 阻塞等待
}

// Stop 停止 tcp 服务器
func (server *TcpServer) Stop() {
	//TODO
}

// SetMidWare 设置 tcp 服务中间件
func (server *TcpServer) SetMidWare(middleware iproxy_net.ITcpBaseMidWare) {
	server.MidWare = middleware
}

// NewTcpServer 创建 TcpServer 实例
func NewTcpServer() iproxy_net.ITcpServer {
	server := TcpServer{
		Proto:   utils.Global.Proto,  // 协议
		HostIp:  utils.Global.HostIp, // 监听 HostIp 的 tcp 连接请求
		Port:    utils.Global.Port,   // 监听的 tcp 端口
		MidWare: nil,
	}
	return &server
}
