package znet

import (
	"errors"
	"fmt"
	"github.com/jodealter/zinx/utils"
	"github.com/jodealter/zinx/ziface"
	"io"
	"net"
)

type Connection struct {
	//当前connection隶属于那个server
	TcpServer ziface.IServer

	//当前链接的socket套接字
	Conn *net.TCPConn

	//链接ID
	ConnID uint32

	//当前链接状态
	isClose bool

	//告知当前链接已经退出 的channel 由Reader告知Writer
	ExitChan chan bool

	//创建一个无缓冲的通道 用于读写之间的通道
	MsgChan chan []byte
	//当前链接处理的rouder
	//同sercer的那个

	//处理msgid 与 api的关系
	MsgHandler ziface.IMsgHandler
}

func (c *Connection) StartReader() {
	fmt.Println("[Reader goroutine is running]")

	defer fmt.Println("connid ", c.ConnID, "[Reader is exit],remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {

		//读取最大512字节的数据到buf中
		/*
			buf := make([]byte, utils.GlobalObject.MaxPackageSize)
			_, err := c.Conn.Read(buf)
			if err != nil {
				continue
			}
		*/
		//创建一个拆包对象
		dp := NewDataPack()

		//读取客户端的Msg Head 二进制流 8字节的数据
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), headData)
		if err != nil {
			fmt.Println("read msghead error : ", err)
			break
		}
		//拆包，得到MsgId与吗、MsgDataLen 放在msg消息中
		msg, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("unpack error :", err)
			break
		}

		//再次读取数据存放在data 中
		var data []byte
		if msg.GetMsglen() > 0 {
			data = make([]byte, msg.GetMsglen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg error ：", err)
				break
			}
		}
		msg.SetData(data)
		//得到当前connect的请求(request)
		request := Request{
			msg:  msg,
			conn: c,
		}

		if utils.GlobalObject.WorkPoolSize > 0 {
			//已经开启了工作池机制，将请求发送给worker就好
			c.MsgHandler.SendMsgToTaskQueue(&request)
		} else {
			//未开启线程池机制，直接开go程
			go c.MsgHandler.DoMsgHandler(&request)
		}
	}
}

/*
写消息的gorotinue 专门发送给客户的模块
*/
func (c *Connection) StartWrite() {
	fmt.Println("[Writer gorotinue is running...]")
	defer fmt.Println(c.Conn.RemoteAddr().String(), "  [conn Writer exit!]")
	for {
		select {
		case data := <-c.MsgChan:
			{
				if _, err := c.Conn.Write(data); err != nil {
					fmt.Println("Send data error,", err)
					return
				}
			}
		case <-c.ExitChan:
			//代表reader 已经退出
			return
		}
	}
}

// 提供sendmsg方法
func (c *Connection) SendMsg(msgid uint32, data []byte) error {
	if c.isClose == true {
		return errors.New("Connection closed when sendmsg")
	}
	//将message进行封包
	dp := DataPack{}
	binaryMsg, err := dp.Pack(NewMessage(msgid, data))
	if err != nil {
		fmt.Println("pack error msgid =  ", msgid)
		return errors.New("Pack error msg")
	}

	//将数据发送给客户端
	c.MsgChan <- binaryMsg
	return nil
}
func (c *Connection) Start() {
	fmt.Println("conn is star....   connid = ", c.ConnID)

	//开启读写的协程
	go c.StartReader()
	go c.StartWrite()
}

func (c *Connection) Stop() {
	fmt.Println("connect stop().. connid = ", c.ConnID)
	if c.isClose == true {
		return
	}

	c.isClose = true

	//关闭链接
	c.Conn.Close()

	//告知Writer关闭
	c.ExitChan <- true

	//将当前链接从ConnMgr中删除
	c.TcpServer.GetConnMgr().Remove(c)
	//回收资源
	close(c.ExitChan)
	close(c.MsgChan)
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

func NewConnect(server ziface.IServer, conn *net.TCPConn, connID uint32, handler ziface.IMsgHandler) *Connection {
	c := &Connection{
		TcpServer:  server,
		Conn:       conn,
		ConnID:     connID,
		MsgHandler: handler,
		isClose:    false,
		MsgChan:    make(chan []byte),
		ExitChan:   make(chan bool, 1),
	}
	//将conn加入到connmanager中
	c.TcpServer.GetConnMgr().Add(c)

	return c
}
