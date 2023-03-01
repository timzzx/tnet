# tcp长连接

server启动
```go
package main

import (
	"github/timzzx/tnet"

	"github/timzzx/tnet/handlers"
)

func main() {
    // 初始化server
	s := tnet.NewServer()
    //上面自定义的handler
	h := handlers.NewTestHandler(1)
    // 注册到服务中
	s.AddHandlers(1, h)
    // 服务启动
	s.Start()
}


```
client启动

```go
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
		// 发送消息 这里的1就是路由id.
        msg, err := tnet.PackSend(1, "test", conn)
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

```

## Demo
[地址](https://github.com/timzzx/tnet-chat)

## 总结

写的很粗糙，目的是实现根据路由id执行自定义的handler

```
.
├── Client
│   └── main.go // 客户端实现
├── LICENSE
├── MsgPack.go
├── README.md
├── Server
│   └── server.go // 服务启动
├── Server.go // 服务实现
├── go.mod
├── go.sum
├── handlers // 自定义handler
│   ├── Test2Handler.go
│   └── TestHandler.go
└── types // handler接口
    └── Handler.go
```