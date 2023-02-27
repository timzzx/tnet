package main

import (
	"github/timzzx/tnet"
	"os"
	"os/signal"
	"syscall"

	"github/timzzx/tnet/handlers"
)

func main() {
	s := tnet.NewServer()
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		select {
		case <-c:
			{
				s.Stop() // 服务器退出
				os.Exit(1)
			}
		}
	}()

	h := handlers.NewTestHandler(1)
	s.AddHandlers(1, h)
	h2 := handlers.NewTestHandler(2)
	s.AddHandlers(2, h2)
	s.Start()
}
