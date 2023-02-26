package main

import (
	"github/timzzx/tnet"

	"github/timzzx/tnet/handlers"
)

func main() {
	s := tnet.NewServer()
	h := handlers.NewTestHandler(1)
	s.AddHandlers(1, h)
	h2 := handlers.NewTestHandler(2)
	s.AddHandlers(2, h2)
	s.Start()
}
