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
	// IP 版本
	IPVersion string
	// 监听的 IP 地址
	IP string
	// 监听的端口
	Port int
}

// Start 启动服务器
func (server *Server) Start() {
	log.Printf("Start server %v, IP %v, port %v\n", server.Name, server.IP, server.Port)
	go func() { //! tcpListener 协程

		//! 解析 TCP 地址
		tcpAddr, err := net.ResolveTCPAddr(server.IPVersion, fmt.Sprintf("%v:%v", server.IP, server.Port))
		if err != nil {
			log.Fatalln("Resolve TCP addr err", err.Error()) //! followed by a call to os.Exit(1).
		}

		//! 监听所有 IP 地址的连接请求
		tcpListener, err := net.ListenTCP(server.IPVersion, tcpAddr)
		if err != nil {
			log.Fatalln("Listen TCP err", err.Error())
		}
		log.Printf("Start server %v ok\n", server.Name)

		//! 阻塞等待客户端的连接请求
		for {
			conn, err := tcpListener.Accept() // 收到客户端的连接请求
			if err != nil {
				log.Println("Accept err", err.Error())
				continue
			}
			// 与客户端建立 TCP 连接
			go func() { //! IO 协程
				for {
					buf := make([]byte, 512)
					readBytes, err := conn.Read(buf)
					if err != nil {
						log.Println("Read err", err.Error())
						continue
					}
					fmt.Printf("Read %v bytes: %v\n", readBytes, string(buf[:readBytes]))
					// 回声 echo 服务器
					if _ /* writeBytes */, err := conn.Write(buf[:readBytes]); err != nil {
						log.Println("Write err", err.Error())
						continue
					}
				}
			}()
		}
	}()
}

// Serve 运行 TCP 服务
func (server *Server) Serve() {
	//! 启动服务器，运行 TCP 服务
	server.Start()
	//TODO

	// 阻塞等待
	select {}
}

// Stop 停止服务器
func (server *Server) Stop() {
	//TODO
}

// NewServer 创建 Server 实例
func NewServer(name string) proxy_inet.IServer {
	server := Server{
		Name:      name,      // 服务器名
		IPVersion: "tcp4",    // IP 版本
		IP:        "0.0.0.0", // 监听所有 IP 地址的连接请求
		Port:      3333,      // 监听的端口
	}
	return &server
}
