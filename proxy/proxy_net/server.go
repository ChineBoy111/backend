package proxy_net

import (
	"fmt"
	"log"
	"net"

	"bronya.com/proxy/proxy_inet"
)

// Server 实现 IServer 接口
type Server struct {
	// 服务器名
	Name string
	// IP 版本
	IPVersion string
	// 服务器监听的 IP 地址
	IP string
	// 服务器监听的端口
	Port int
}

// Start 启动服务器
func (server *Server) Start() {
	log.Printf("[Start] Server IP %v, Port %v\n", server.IP, server.Port)
	tcpAddr, err := net.ResolveTCPAddr(server.IPVersion, fmt.Sprintf("%s:%d", server.IP, server.Port))
	if err != nil {
		log.Fatalln("Resolve TCP address error") //! followed by a call to os.Exit(1).
	}
	net.ListenTCP(server.IPVersion, tcpAddr)
	
}

// Stop 停止服务器
func (server *Server) Stop() {

}

// Server 运行 TCP 服务
func (server *Server) Serve() {
	server.Start()
}

// 创建 Server 实例
func NewServer(name string) proxy_inet.IServer {
	server := Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0", //! 监听所有 IP 地址的连接请求
		Port:      3333,
	}
	return &server
}
