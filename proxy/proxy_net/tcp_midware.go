package proxy_net

import (
	"bronya.com/proxy/iproxy_net"
)

type TcpBaseMidware struct {
}

// PreHandler Handler 前的 hook 方法
func (midware *TcpBaseMidware) PreHandler(iproxy_net.ITcpReq) {
}

// Handler 处理收到的 tcp 数据
func (midware *TcpBaseMidware) Handler(iproxy_net.ITcpReq) {
}

// PostHandler Handler 后的 hook 方法
func (midware *TcpBaseMidware) PostHandler(iproxy_net.ITcpReq) {
}
