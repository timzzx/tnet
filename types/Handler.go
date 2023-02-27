package types

type Handler interface {
	Do([]byte, Connection)
}
