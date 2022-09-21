package net

import (
	"github.com/lovelacelee/clsgo/v1/net/tcp"
)

// Wrappers of tcp
type Conn = tcp.Conn
type Retry = tcp.Retry

// User defined TCP protocol
type TcpProtocol interface {
	// Specify the server name
	ServerName() string
	// Return a new object or the exist
	Instance() TcpProtocol
	// Recv data from peer
	// Return proccess result
	HandleMessage(conn *Conn) ([]byte, error)
}
