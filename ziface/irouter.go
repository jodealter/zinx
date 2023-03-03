package ziface

/*
	路由抽象接口
	路由里的数据都是request
*/

type IRouter interface {

	// PreHandle 处理conn业务之前的钩子方法
	PreHandle(request IRequest)

	// Handle 处理conn时的业务方法
	Handle(request IRequest)

	// PostHandle 处理conn业务之后的方法
	PostHandle(request IRequest)
}
