package utils

import (
	"encoding/json"
	"github.com/jodealter/zinx/ziface"
	"io/ioutil"
)

/*
客户构建一个服务器，通过这个进行配置
*/
type GlobalObj struct {
	/*
		server的配置
	*/
	TcpServer ziface.IServer //当前zinx全局的Server对象
	Host      string         //监听的ip
	TcpPort   int            //监听的端口
	Name      string         //服务器的名字

	/*zinx*/
	Version          string //zinx的版本
	MaxConn          int    //允许链接的最大数木
	MaxPackageSize   uint32 //数据包的最大值
	WorkPoolSize     uint32 //当前业务工作worker的goroutine的数量
	MaxWorkerTaskLen uint32 //队列长度
}

/*定义一个全局的Globalobj*/

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

var GlobalObject *GlobalObj

func init() {
	GlobalObject = &GlobalObj{
		Name:           "ZinxServerapp",
		Version:        "v0.9",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,

		WorkPoolSize:     10,
		MaxWorkerTaskLen: 1024,
	}
	GlobalObject.Reload()
}
