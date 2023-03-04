package ziface

/*
封包拆包处理
直接面向tcp数据流，用于处理粘包问题
*/
type IDataPack interface {
	//获取包的头部
	GetHeadLen() uint32

	//封包
	Pack(msg IMessage) ([]byte, error)

	//拆包
	UnPack([]byte) (IMessage, error)
}
