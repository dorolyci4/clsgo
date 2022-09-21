// Package net provides basic network functions,
// TCP/UDP/WebSocket quick implements.
package net

import (
	"strconv"
	"time"

	"github.com/lovelacelee/clsgo/v1/log"
	"github.com/lovelacelee/clsgo/v1/net/tcp"
)

type TcpServer = *tcp.Server

// Close the TCP connection if any error occurred or the connection lifetime ends.
func connectionClose(conn *tcp.Conn) {
	conn.Close()
}

// NewTcpServer provides RECV-HANDLE-SENDBACK routine of TCP server module,
// Keepalive message must be handled carefully, because all communication
// data packets are walk on the single connection instance.
// NewTcpServer will be block running after Run() called.
// Use TcpServer.Close() to terminate the server.
func NewTcpServer(host string, port int, tcpreuse bool, proto TcpProtocol) TcpServer {
	log.Debugfi("TCP server[%v] listen %s:%d", proto.ServerName(), host, port)
	return tcp.NewServer(host+":"+strconv.Itoa(port), tcpreuse, func(conn *tcp.Conn) {
		log.Debugfi("New connection [%v] %v", proto.ServerName(), conn.RemoteAddr().String())
		defer connectionClose(conn)
		p := proto.Instance()
		for {
			data, err := p.HandleMessage(conn)
			if err != nil {
				break
			}
			err = conn.Send(data, tcp.Retry{Count: 3, Interval: time.Microsecond})
			if err != nil {
				break
			}
		}
		log.Debugfi("Connection %v with[%v] closed. ", conn.RemoteAddr().String(), proto.ServerName())
	}, proto.ServerName())
}
