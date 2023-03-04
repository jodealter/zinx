package znet

import (
	"fmt"
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
	Router ziface.IRouter
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
}

func (s *Server) Start() {
	fmt.Printf("[strat] server listenr at ip:%s, port:%d\n", s.IP, s.Port)
	go func() {
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
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept err ", err)
				continue
			}
			var cid uint32
			cid = 0

			//奖处理新连接的业务方法和conn进行绑定
			dealconn := NewConnect(conn, cid, s.Router)
			cid++

			go dealconn.Start()
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()

	//TODO 做一些其他事情

	select {}
}
func NewServr(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		Router:    nil,
	}
	return s
}
