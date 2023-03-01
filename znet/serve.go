package znet

import "github.com/jodealter/zinx/ziface"

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

func (s *Server) Start() {

}

func (s *Server) Stop() {

}

func (s *Server) Serve() {

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
