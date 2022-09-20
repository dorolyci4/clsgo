package net

import (
	"github.com/gogf/gf/v2/net/gtcp"
)

// Wrappers of gtcp
type Conn = gtcp.Conn
type Retry = gtcp.Retry

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
