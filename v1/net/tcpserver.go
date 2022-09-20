// Package net provides basic network functions,
// TCP/UDP/WebSocket quick implements.
package net

import (
	"strconv"
	"time"

	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/lovelacelee/clsgo/v1/log"
)

type TcpServer = *gtcp.Server

// Close the TCP connection if any error occurred or the connection lifetime ends.
func connectionClose(conn *gtcp.Conn) {
	conn.Close()
}

// NewTcpServer provides RECV-HANDLE-SENDBACK routine of TCP server module,
// Keepalive message must be handled carefully, because all communication
// data packets are walk on the single connection instance.
// NewTcpServer will be block running after Run() called.
// Use TcpServer.Close() to terminate the server.
func NewTcpServer(host string, port int, proto TcpProtocol) TcpServer {
	log.Debugfi("TCP server[%v] listen %s:%d", proto.ServerName(), host, port)
	return gtcp.NewServer(host+":"+strconv.Itoa(port), func(conn *gtcp.Conn) {
		log.Debugfi("New connection [%v] %v", proto.ServerName(), conn.RemoteAddr().String())
		defer connectionClose(conn)
		p := proto.Instance()
		for {
			data, err := p.HandleMessage(conn)
			if err != nil {
				break
			}
			err = conn.Send(data, gtcp.Retry{Count: 3, Interval: time.Microsecond})
			if err != nil {
				break
			}
		}
		log.Debugfi("Connection %v with[%v] closed. ", conn.RemoteAddr().String(), proto.ServerName())
	}, proto.ServerName())
}
