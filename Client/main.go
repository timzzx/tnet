package main

import (
	"fmt"
	"github/timzzx/tnet"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.13:9999")
	if err != nil {
		fmt.Println("连接失败", err)
	}
	defer conn.Close()
	for {
		// 发送消息
		msg, err := tnet.Pack(1, []byte("test"))
		conn.Write(msg)
		if err != nil {
			fmt.Println("消息发送失败", err)
			return
		}

		// 接收消息
		_, data, err := tnet.Unpack(conn)
		if err != nil {
			fmt.Println("消息收回", err)
			conn.Close()
			return
		}

		fmt.Println(string(data))
		time.Sleep(time.Second)
	}
}
