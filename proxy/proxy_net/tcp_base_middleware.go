package proxy_net

import (
	"bronya.com/proxy/iproxy_net"
)

type TcpBaseMiddleware struct {
}

// PreHandler PacketHandler 前的 hook 方法
func (middleware *TcpBaseMiddleware) PreHandler(request iproxy_net.ITcpRequest) {
}

// PacketHandler 处理收到的 tcp 数据包
func (middleware *TcpBaseMiddleware) PacketHandler(request iproxy_net.ITcpRequest) {
}

// PostHandler PacketHandler 后的 hook 方法
func (middleware *TcpBaseMiddleware) PostHandler(request iproxy_net.ITcpRequest) {
}
