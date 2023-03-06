package znet

import (
	"errors"
	"fmt"
	"github.com/jodealter/zinx/ziface"
	"sync"
)

type ConnManager struct {
	connections map[uint32]ziface.IConnection
	connLock    sync.RWMutex
}

func (c *ConnManager) Add(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	c.connections[conn.GetConnID()] = conn
	fmt.Println("onnection = ", conn.GetConnID(), " Add to Connmanager successfully ，conn num = ", c.Len())
}

func (c *ConnManager) Remove(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	delete(c.connections, conn.GetConnID())
	fmt.Println("onnection = ", conn.GetConnID(), " Remove from Connmanager successfully ，conn num = ", c.Len())

}

func (c *ConnManager) Get(ConnID uint32) (ziface.IConnection, error) {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	if conn, ok := c.connections[ConnID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection Not Found!!!")
	}
}

func (c *ConnManager) Len() int {
	return len(c.connections)
}

func (c *ConnManager) ClearConn() {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	for connID, conn := range c.connections {
		conn.Stop()
		delete(c.connections, connID)
	}
	fmt.Println("Clear All connections succ!")
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}
