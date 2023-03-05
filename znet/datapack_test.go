package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {
	/*
		模拟一个服务器
	*/

	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listener error : ", err)
		return
	}

	//
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("server accept error : ", err)
				return
			}
			go func(conn net.Conn) {
				//处理客户请求
				//拆包过程
				//定义一个拆包对象
				dp := NewDataPack()
				for {
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head error :", err)
						return
					}
					msgHead, err := dp.UnPack(headData)
					if err != nil {
						fmt.Println("server unpack error : ", err)
						return
					}
					if msgHead.GetMsglen() > 0 {
						//说明有数据，可以进行二次读取
						msg := msgHead.(*Message)
						msg.data = make([]byte, msg.GetMsglen())
						_, err := io.ReadFull(conn, msg.data)
						if err != nil {
							fmt.Println("server second read error : ", err)
							return
						}
						fmt.Println("recv Msgid :", msg.Id, " datalen :", msg.Datalen, " data:", string(msg.data))
					}
				}
			}(conn)
		}
	}()

	/*
		模拟客户端
	*/
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client error ", err)
		return
	}
	//创建一个封包对象
	dp := NewDataPack()

	//封装两个包
	msg1 := &Message{
		Id:      1,
		Datalen: 4,
		data:    []byte{'z', 'i', 'n', 'x'},
	}
	senddata1, err := dp.Pack(msg1)
	if err != nil {
		return
	}
	//第二个包
	msg2 := &Message{
		Id:      2,
		Datalen: 4,
		data:    []byte{'n', 'i', 'h', 'a'},
	}
	senddata2, err := dp.Pack(msg2)
	if err != nil {
		return
	}
	//将2个包连在一起
	senddata := append(senddata1, senddata2...)
	conn.Write(senddata)

	select {}
}
