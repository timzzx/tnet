package tnet

import (
	"fmt"
	"github/timzzx/tnet/types"
	"net"
	"sync"
)

type Server struct {
	Name        string
	Port        string
	Connections map[string]net.Conn
	Handlers    map[int]types.Handler

	mu sync.Mutex
}

func NewServer() *Server {
	return &Server{
		Name:        "tcp",
		Port:        "9999",
		Connections: make(map[string]net.Conn),
		Handlers:    make(map[int]types.Handler),
	}
}

func (s *Server) Start() {
	listen, err := net.Listen("tcp", "0.0.0.0:"+s.Port)
	if err != nil {
		fmt.Println("监听失败：", err)
		return
	}

	fmt.Println("TCP服务启动成功...")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("建立连接失败", err)
			continue
		}
		fmt.Println("连接建立成功")

		// 获取用户id
		uid := conn.RemoteAddr().String()
		// 连接加入全局
		s.addConnections(uid, conn)

		// 逻辑控制
		go s.proceess(conn)

	}
}

// 逻辑控制
func (s *Server) proceess(conn net.Conn) {
	defer conn.Close()
	for {
		// 获取消息
		routerID, data, err := Unpack(conn)
		if err != nil {
			fmt.Println("消息解析失败：", err)
			return
		}

		// 根据路由id调用处理逻辑
		s.doHandler(routerID, data, conn)
	}
}

// 根据路由调用处理逻辑
func (s *Server) doHandler(id int, data string, conn net.Conn) {
	s.Handlers[id].Do(data, conn)
}

// 连接加入全局
func (s *Server) addConnections(uid string, conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Connections[uid] = conn

}

// 添加handler
func (s *Server) AddHandlers(id int, handler types.Handler) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Handlers[id] = handler
}

func (s *Server) Stop() {

}
