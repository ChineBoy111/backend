package proxy_net

import (
	"bronya.com/proxy/iproxy_net"
)

type TcpBaseMidWare struct {
}

// PreHandler Handler 前的 hook 方法
func (midWare *TcpBaseMidWare) PreHandler(iproxy_net.ITcpReq) {
}

// Handler 处理收到的 tcp 数据
func (midWare *TcpBaseMidWare) Handler(iproxy_net.ITcpReq) {
}

// PostHandler Handler 后的 hook 方法
func (midWare *TcpBaseMidWare) PostHandler(iproxy_net.ITcpReq) {
}
