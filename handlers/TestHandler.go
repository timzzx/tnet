package handlers

import (
	"github.com/timzzx/tnet"
	"github.com/timzzx/tnet/types"
)

type TestHandler struct {
	id int
}

func NewTestHandler(id int) types.Handler {
	return &TestHandler{id: id}
}

func (h *TestHandler) Do(data []byte, agent types.Connection) {

	// fmt.Println("handlerID:", h.id, "消息:", )
	// 封包并发送
	msg, _ := tnet.Pack(h.id, data)
	agent.Send(msg)
	agent.Cancel()
}
