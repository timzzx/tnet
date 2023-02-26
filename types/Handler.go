package types

import "net"

type Handler interface {
	Do(string, net.Conn)
}
