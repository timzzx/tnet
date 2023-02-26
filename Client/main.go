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
		msg, err := tnet.PackSend(1, "test", conn)
		conn.Write(msg)
		if err != nil {
			fmt.Println("消息发送失败", err)
			return
		}
		msg, err = tnet.PackSend(2, "test2", conn)
		conn.Write(msg)
		if err != nil {
			fmt.Println("消息发送失败", err)
			return
		}

		// 接收消息
		_, data, err := tnet.Unpack(conn)
		if err != nil {
			fmt.Println("消息收回", err)
			return
		}

		fmt.Println(data)
		time.Sleep(time.Second)
	}
}
