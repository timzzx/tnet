package handlers

import (
	"github/timzzx/tnet"
	"github/timzzx/tnet/types"
	"net"
)

type TestHandler struct {
	id int
}

func NewTestHandler(id int) types.Handler {
	return &TestHandler{id: id}
}

func (h *TestHandler) Do(data string, conn net.Conn) {

	// fmt.Println("handlerID:", h.id, "消息:", )
	// 封包并发送
	msg, _ := tnet.PackSend(h.id, "handler发送:"+data, conn)
	conn.Write(msg)
}
