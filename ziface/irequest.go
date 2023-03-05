package ziface

// 这是ziface
type IRequest interface {
	GetConnection() IConnection
	GetData() []byte
	GetMsgId() uint32
}
