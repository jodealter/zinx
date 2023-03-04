package znet

type Message struct {
	Id      uint32 //消息的id
	Datalen uint32 //消息的长度
	data    []byte //消息的内容
}

func (m *Message) GetMsgId() uint32 {
	return m.Id
}

func (m *Message) SetMsgId(u uint32) {
	m.Id = u
}

func (m *Message) GetMsglen() uint32 {
	return m.Datalen
}

func (m *Message) SetDatalen(u uint32) {
	m.Datalen = u
}

func (m *Message) GetData() []byte {
	return m.data
}

func (m *Message) SetData(bytes []byte) {
	m.data = bytes
}
