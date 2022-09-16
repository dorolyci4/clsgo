package net

import (
	"github.com/gogf/gf/v2/net/gtcp"
)

// Wrappers of gtcp
type Conn = gtcp.Conn
type Retry = gtcp.Retry

// User defined TCP protocol
type TcpProtocol interface {
	// Method of handling message
	HandleMessage(conn *Conn) ([]byte, error)
	Instance() TcpProtocol

	// Send methods
	Login(conn *Conn) error
	KeepAlive(conn *Conn) error
}
