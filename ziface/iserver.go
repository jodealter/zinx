package ziface

type IServer interface {
	//启动服务器
	Start()

	//停止服务器
	Stop()

	//运行服务器
	Serve()

	//添加router的方法
	AddRouter(msgID uint32, router IRouter)

	//返回ConnManager
	GetConnMgr() IConnManager
	//注册OnConnStart钩子方法
	SetOnConnStart(func(conn IConnection))
	//注册OnConnStop钩子方法
	SetOnConnStop(func(conn IConnection))
	//调用OnConnStart函数的方法
	CallOnConnStart(conn IConnection)
	//调用OnConnStop函数的方法
	CallOnConnStop(conn IConnection)
}
