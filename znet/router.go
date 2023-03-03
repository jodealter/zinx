package znet

import "github.com/jodealter/zinx/ziface"

// BaseRouter 这样可以做到，当使用route时，可以先继承BaseRouter，然后选择适当的方法进行继承，例如舍去Pre与Post，只保留Handle
type BaseRouter struct {
}

func (b *BaseRouter) PreHandle(request ziface.IRequest) {
}

func (b *BaseRouter) Handle(request ziface.IRequest) {
}

func (b *BaseRouter) PostHandle(request ziface.IRequest) {
}
