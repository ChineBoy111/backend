package proxy_net

import (
	"bronya.com/proxy/iproxy_net"
)

type TcpMiddleware struct {
}

// PreHandler PacketHandler 前的 hook 方法
func (middleware *TcpMiddleware) PreHandler(iproxy_net.ITcpRequest) {
}

// PacketHandler 处理收到的 tcp 数据包
func (middleware *TcpMiddleware) PacketHandler(iproxy_net.ITcpRequest) {
}

// PostHandler PacketHandler 后的 hook 方法
func (middleware *TcpMiddleware) PostHandler(iproxy_net.ITcpRequest) {
}
