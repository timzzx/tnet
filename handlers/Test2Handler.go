package handlers

import (
	"github/timzzx/tnet"
	"github/timzzx/tnet/types"
	"net"
)

type Test2Handler struct {
	id int
}

func NewTest2Handler(id int) types.Handler {
	return &TestHandler{id: id}
}

func (h *Test2Handler) Do(data string, conn net.Conn) {

	// fmt.Println("handlerID:", h.id, "消息:", )
	// 封包并发送
	msg, _ := tnet.PackSend(h.id, "handler2发送:"+data, conn)
	conn.Write(msg)
}
