package network

import "bronya.com/net-proxy/inetwork"

// TcpReq 封装 tcp 连接和 tcp 消息的 tcp 请求
type TcpReq struct {
	Conn inetwork.ITcpConn // tcp 连接
	Msg  inetwork.ITcpMsg  // tcp 消息
}

// GetConn 获取 tcp 连接
func (req *TcpReq) GetConn() inetwork.ITcpConn {
	return req.Conn
}

// GetMsgLen 获取 tcp 消息的长度
func (req *TcpReq) GetMsgLen() uint32 {
	return req.Msg.GetLen()
}

// GetMsgId 获取 tcp 消息的 id
func (req *TcpReq) GetMsgId() uint32 {
	return req.Msg.GetId()
}

// GetMsgData 获取 tcp 消息的数据
func (req *TcpReq) GetMsgData() []byte {
	return req.Msg.GetData()
}
