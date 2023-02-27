package tnet

import (
	"context"
	"github/timzzx/tnet/types"
	"net"
)

type Connection struct {
	Uid    string
	Conn   net.Conn
	ctx    context.Context
	cancel context.CancelFunc
}

func NewConnection(uid string, conn net.Conn, ctx context.Context, cancel context.CancelFunc) types.Connection {
	return &Connection{
		Uid:    uid,
		Conn:   conn,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (c *Connection) GetConn() net.Conn {
	return c.Conn
}

func (c *Connection) GetUid() string {
	return c.Uid
}

func (c *Connection) Send(data []byte) (n int, err error) {
	return c.Conn.Write(data)
}

func (c *Connection) Cancel() {
	c.cancel()
}

func (c *Connection) Ctx() context.Context {
	return c.ctx
}
