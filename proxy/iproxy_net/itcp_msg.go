package iproxy_net

type ITcpMsg interface {
	// GetMsgLen 获取 tcp 消息长度
	GetMsgLen() uint32

	// GetMsgId 获取 tcp 消息 id
	GetMsgId() uint32

	// GetMsgData 获取 tcp 消息数据
	GetMsgData() []byte

	// SetMsgLen 设置 tcp 消息长度
	SetMsgLen(uint32)

	// SetMsgId 设置 tcp 消息 id
	SetMsgId(uint32)

	// SetMsgData 设置 tcp 消息数据
	SetMsgData([]byte)
}
