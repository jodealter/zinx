package ziface

type IMessage interface {

	//id的读与设置
	GetMsgId() uint32
	SetMsgId(uint32)

	//消息长度的读与设置
	GetMsglen() uint32
	SetDatalen(uint32)

	//消息内容的读与设置
	GetData() []byte
	SetData([]byte)
}
