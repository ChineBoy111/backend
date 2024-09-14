package proxy_net

import "bronya.com/proxy/iproxy_net"

type TcpMiddleware struct {
}

// PreHandler TcpConnector.packetHandler 前的 hook 方法
func (tcpMiddleware *TcpMiddleware) PreHandler(tcpReq iproxy_net.ITcpRequest) {
}

func (tcpMiddleware *TcpMiddleware) Handler(tcpReq iproxy_net.ITcpRequest) {
}

// PostHandler TcpConnector.packetHandler 后的 hook 方法
func (tcpMiddleware *TcpMiddleware) PostHandler(tcpReq iproxy_net.ITcpRequest) {
}
