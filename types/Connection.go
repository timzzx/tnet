package types

import (
	"context"
	"net"
)

type Connection interface {
	GetConn() net.Conn
	GetUid() string
	Send(data []byte) (n int, err error)
	Cancel()
	Ctx() context.Context
}
