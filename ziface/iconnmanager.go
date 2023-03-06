package ziface

type IConnManager interface {

	//添加链接
	Add(conn IConnection)

	//删除链接
	Remove(conn IConnection)

	//跟据connID获取链接
	Get(ConnID uint32) (IConnection, error)

	//返回链接总数
	Len() int

	//清楚全部连接
	ClearConn()
}
