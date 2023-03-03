package znet

import (
	"errors"
	"fmt"
	"github.com/jodealter/zinx/ziface"
	"net"
)

type Server struct {

	//服务器名称
	Name string

	//服务器绑定的版本
	IPVersion string

	//服务器监听的ip
	IP string

	//服务器监听的端口
	Port int
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	//回显业务
	fmt.Println("[Conn handle]callbacktoclient...")
	if _, err := conn.Write(data); err != nil {
		fmt.Println("write to client error", err)
		return errors.New("callbacktoclient error")
	}
	cnt++
	return nil
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
			dealconn := NewConnect(conn, cid, CallBackToClient)
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
	}
	return s
}
