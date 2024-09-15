package proxy_net

type TcpMsg struct {
	MsgLen  uint32
	MsgId   uint32
	MsgData []byte
}

// GetMsgLen 获取 tcp 消息长度
func (msg *TcpMsg) GetMsgLen() uint32 {
	return msg.MsgLen
}

// GetMsgId 获取 tcp 消息 id
func (msg *TcpMsg) GetMsgId() uint32 {
	return msg.MsgId
}

// GetMsgData 获取 tcp 消息数据
func (msg *TcpMsg) GetMsgData() []byte {
	return msg.MsgData
}

// SetMsgLen 设置 tcp 消息长度
func (msg *TcpMsg) SetMsgLen(msgLen uint32) {
	msg.MsgLen = msgLen
}

// SetMsgId 设置 tcp 消息 id
func (msg *TcpMsg) SetMsgId(msgId uint32) {
	msg.MsgId = msgId
}

// SetMsgData 设置 tcp 消息数据
func (msg *TcpMsg) SetMsgData(msgData []byte) {
	msg.MsgData = msgData
}
