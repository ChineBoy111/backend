package proxy_net

import (
	"bronya.com/proxy/iproxy_net"
	"fmt"
	"log"
	"net"
)

// TcpServer 实现 ITcpServer 接口
type TcpServer struct {
	// tcp 服务器名
	Name string
	// ip 版本
	IpVer string
	// 监听的 ip 地址
	Ip string
	// 监听的端口
	Port int
}

// Start 启动 tcp 服务器
func (tcpServer *TcpServer) Start() {
	log.Printf("Start tcpServer %v, ip %v, port %v\n", tcpServer.Name, tcpServer.Ip, tcpServer.Port)
	go func() { //! 负责监听所有 ip 地址的 tcp 连接请求的 goroutine

		//! 解析 tcp 地址
		tcpAddr, err := net.ResolveTCPAddr(tcpServer.IpVer, fmt.Sprintf("%v:%v", tcpServer.Ip, tcpServer.Port))
		if err != nil {
			log.Println("ResolveTCPAddr err", err.Error())
			return
		}

		//! 监听所有 ip 地址的 tcp 连接请求
		tcpListener, err := net.ListenTCP(tcpServer.IpVer, tcpAddr)
		if err != nil {
			log.Println("ListenTCP err", err.Error())
			return
		}
		log.Printf("Start tcpServer %v ok, listening %v:%v\n", tcpServer.Name, tcpServer.Ip, tcpServer.Port)

		var connId uint32 = 0
		//! 阻塞等待客户端的 tcp 连接请求
		for {
			socket, err := tcpListener.AcceptTCP() // 收到客户端的 tcp 连接请求
			if err != nil {
				log.Println("AcceptTCP err", err.Error())
				continue
			}

			//! 已建立连接的 tcpConn
			tcpConn := NewTcpConn(socket, connId, DataHandler)
			connId++
			go tcpConn.Start() //! 负责处理 tcp 连接的 goroutine
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

// Serve 运行 tcp 服务
func (tcpServer *TcpServer) Serve() {
	//! 启动 tcp 服务器，运行服务
	tcpServer.Start()
	//TODO

	//! ==================== 阻塞等待 ====================
	select {}
}

// Stop 停止 tcp 服务器
func (tcpServer *TcpServer) Stop() {
	//TODO
}

// NewServer 创建 TcpServer 实例
func NewServer(name string) iproxy_net.ITcpServer {
	tcpServer := TcpServer{
		Name:  name,      // 服务器名
		IpVer: "tcp4",    // ip 版本
		Ip:    "0.0.0.0", // 监听所有 ip 地址的 tcp 连接请求
		Port:  3333,      // 监听的端口
	}
	return &tcpServer
}
