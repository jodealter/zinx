package ziface

type IMsgHandler interface {
	//实现两个方法
	//第一个添加具体的处理逻辑

	AddRouter(msgID uint32, router IRouter)
	//第二个是进行调用
	DoMsgHandler(request IRequest)
}
