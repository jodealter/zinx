package ziface

type IMsgHandler interface {
	//实现两个方法
	//第一个添加具体的处理逻辑

	AddRouter(msgID uint32, router IRouter)
	//第二个是进行调用
	DoMsgHandler(request IRequest)

	//打开工作池
	StartWorkerPool()
	//将工作交给工作队列，并有worker处理
	SendMsgToTaskQueue(request IRequest)
}
