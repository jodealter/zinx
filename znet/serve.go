package znet

import (
	"fmt"
	"github.com/jodealter/zinx/utils"
	"github.com/jodealter/zinx/ziface"
	"net"
)

// ceshi
type Server struct {

	//服务器名称
	Name string

	//服务器绑定的版本
	IPVersion string

	//服务器监听的ip
	IP string

	//服务器监听的端口
	Port int

	//添加一个router，也就是这个server绑定的业务
	//本人觉得不太好，应该做一个切片类型的。这样，使用的时候可以绑定多个方法
	MsgHandler ziface.IMsgHandler

	//该server的链接管理模块
	ConnMgr ziface.IConnManager

	//该Server创建链接之后自动调Hook函数OnConnStart
	onConnStart func(conn ziface.IConnection)

	//该Server断开链接之前自动调Hook函数OnConnStop
	onConnStop func(conn ziface.IConnection)
}

func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
	fmt.Printf("add router succ\n")
}

func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name : %s, listener at IP : %s, Port:%d\n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[Zinx] Version : %s , MaxConn : %d , MaxPackageSize : %d\n",
		utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)
	go func() {

		//开启线程池
		s.MsgHandler.StartWorkerPool()
		//获取一个tcp的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err", err)
			return
		}

		//监听
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, "err", err)
			return
		}
		fmt.Println("Start Zinx server succ ,", s.Name, " is Listining...")
		var cid uint32
		cid = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept err ", err)
				continue
			}

			//判断当前链接是否到达上限
			//这里之所以不放在accept之前进行判断，是因为可以扩展功能，比如反馈给客户信息

			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				fmt.Println("too many Connections")
				conn.Close()
				continue
			}
			//奖处理新连接的业务方法和conn进行绑定
			dealconn := NewConnect(s, conn, cid, s.MsgHandler)
			cid++

			go dealconn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] zinx server name ", s.Name)
	s.ConnMgr.ClearConn()
}

func (s *Server) Serve() {
	s.Start()

	//TODO 做一些其他事情

	select {}
}

// 返回connManager的方法
func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

func NewServr(name string) ziface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandler(),
		ConnMgr:    NewConnManager(),
	}
	return s
}

func (s *Server) SetOnConnStart(f func(conn ziface.IConnection)) {
	s.onConnStart = f
}

func (s *Server) SetOnConnStop(f func(conn ziface.IConnection)) {
	s.onConnStop = f
}

func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.onConnStart != nil {
		fmt.Println("----->Call OnConnStart()...")
		s.onConnStart(conn)
	}
}

func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.onConnStop != nil {
		fmt.Println("----->Call OnConnStop()...")
		s.onConnStop(conn)
	}
}
