package znet

import "github.com/jodealter/zinx/ziface"

type Request struct {
	conn ziface.IConnection

	msg ziface.IMessage
}

// 获取该请求的链接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// 获取该请求的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgId() uint32 {
	return r.msg.GetMsgId()
}
