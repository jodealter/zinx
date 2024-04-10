package znet

import (
	"fmt"
	"github.com/jodealter/zinx/utils"
	"github.com/jodealter/zinx/ziface"
	"strconv"
)

//消息处理模块

// 初始化

type MsgHandler struct {

	//处理消息的api
	Apis map[uint32]ziface.IRouter

	//负责处理消息的消息队列
	TaskQuene []chan ziface.IRequest

	//线程池的数量
	WorkPoolSize uint32
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:         map[uint32]ziface.IRouter{},
		WorkPoolSize: utils.GlobalObject.WorkPoolSize,
		TaskQuene:    make([]chan ziface.IRequest, utils.GlobalObject.WorkPoolSize), //工作队列的数量，不是单指一个工作队列
	}
}

func (m *MsgHandler) AddRouter(msgID uint32, router ziface.IRouter) {
	if _, ok := m.Apis[msgID]; ok {
		panic("repeat api ,msgid = " + strconv.Itoa(int(msgID)))
	}
	m.Apis[msgID] = router
	fmt.Println("add api succ msgid = ", msgID, "success")
}

func (m *MsgHandler) DoMsgHandler(request ziface.IRequest) {
	handler, ok := m.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("api msgid is not found ；", request.GetMsgId())

		return
	}

	//根据相应的msgid调度相应的业务
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (m *MsgHandler) StartWorkerPool() {
	//根据workerpoolsize分开启去 worker
	for i := 0; i < int(m.WorkPoolSize); i++ {
		//1.为对应worker的消息队列开辟空间
		m.TaskQuene[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		//2.开启这和worker
		go m.StartOneWorker(i, m.TaskQuene[i])
	}
}

func (m *MsgHandler) StartOneWorker(workerID int, taskquene chan ziface.IRequest) {
	fmt.Println("Worker ID =", workerID, " is  starting")
	for {
		select {
		case request := <-taskquene:
			m.DoMsgHandler(request)
		}
	}
}
func (m *MsgHandler) SendMsgToTaskQueue(request ziface.IRequest) {
	//1.将消息平均分配给不同的worker
	//简单策略，根据客户端id 进行分配
	workID := request.GetConnection().GetConnID() % m.WorkPoolSize
	fmt.Println("ADD ConnId = ", request.GetConnection().GetConnID(),
		" request MsgID = ", request.GetMsgId(), " to WorkerId = ", workID)
	//2.将消息发送给对应的worker的TaskQuene即可
	m.TaskQuene[workID] <- request
}
