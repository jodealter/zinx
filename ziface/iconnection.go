package ziface

import (
	"net"
)

type IConnection interface {
	//启动链接
	Start()

	//停止链接
	Stop()

	//获取当前链接的socket
	GetTCPConnection() *net.TCPConn

	//获取当前链接模块的id
	GetConnID() uint32

	//获取远程客户端的状态
	RemoteAddr() net.Addr

	//发送数据
	SendMsg(msgid uint32, data []byte) error

	//设置链接属性
	SetProperty(key string, value interface{})

	//获取链接属性
	GetProperty(key string) (interface{}, error)

	//移除链接属性
	RemoveProperty(key string)
}

// 定义一个处理链接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
