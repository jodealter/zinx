package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/jodealter/zinx/utils"
	"github.com/jodealter/zinx/ziface"
)

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (d *DataPack) GetHeadLen() uint32 {
	//四字节头部（存放长度），四字节类型
	return 8
}

// 封包方法
func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	//创建一个存放byte字节的缓冲
	databuf := bytes.NewBuffer([]byte{})
	//把datalen写入databuf中
	if err := binary.Write(databuf, binary.LittleEndian, msg.GetMsglen()); err != nil {
		return nil, err
	}

	//把msgid写入
	if err := binary.Write(databuf, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	//将data写入
	if err := binary.Write(databuf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return databuf.Bytes(), nil
}

func (d *DataPack) UnPack(binaryData []byte) (ziface.IMessage, error) {
	//创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	//只解压head信息 得到len与id
	msg := &Message{}

	//读len
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Datalen); err != nil {
		return nil, err
	}

	//读id
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	if utils.GlobalObject.MaxPackageSize > 0 && msg.Datalen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large")
	}
	return msg, nil
}
