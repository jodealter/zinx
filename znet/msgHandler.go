package znet

import (
	"fmt"
	"github.com/jodealter/zinx/ziface"
	"strconv"
)

//消息处理模块

// 初始化

type MsgHandler struct {
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis: map[uint32]ziface.IRouter{},
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
	}

	//根据相应的msgid调度相应的业务
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}
