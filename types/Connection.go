package types

import "net"

type Connection interface {
	GetConn() net.Conn
	GetUid() string
	Send(data []byte) (n int, err error)
}
