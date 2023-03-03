package znet

import (
	"fmt"
	"github.com/jodealter/zinx/ziface"
	"net"
)

type Connection struct {
	//当前链接的socket套接字
	Conn *net.TCPConn

	//链接ID
	ConnID uint32

	//当前链接状态
	isClose bool

	//告知当前链接已经退出 的channel
	ExitChan chan bool

	//当前链接处理的rouder
	//同sercer的那个
	Rouder ziface.IRouter
}

func (c *Connection) StartReader() {
	fmt.Println("reader goroutine is running")

	defer fmt.Println("connid ", c.ConnID, "reader is exit,remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {

		//读取最大512字节的数据到buf中
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			continue
		}

		//得到当前connect的请求(request)
		request := Request{
			data: buf,
			conn: c,
		}

		//执行注册的路由方法
		go func(request ziface.IRequest) {
			c.Rouder.PreHandle(request)
			c.Rouder.Handle(request)
			c.Rouder.PostHandle(request)

		}(&request)
	}
}

func (c *Connection) Start() {
	fmt.Println("conn is star....   connid = ", c.ConnID)
	go c.StartReader()
}

func (c *Connection) Stop() {
	fmt.Println("connect stop().. connid = ", c.ConnID)
	if c.isClose == true {
		return
	}

	c.isClose = true

	//关闭链接
	c.Conn.Close()

	//关闭通道
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	//TODO implement me
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	//TODO implement me
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	//TODO implement me
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	//TODO implement me
	return nil
}

func NewConnect(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Rouder:   router,
		isClose:  false,
		ExitChan: make(chan bool, 1),
	}
	return c
}
