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
}
