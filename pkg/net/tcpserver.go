// Package net provides basic network functions,
// TCP/UDP/WebSocket quick implements.
package net

import (
	"strconv"

	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/lovelacelee/clsgo/pkg/log"
)

// Close the TCP connection if any error occurred or the connection lifetime ends.
func connectionClose(conn *gtcp.Conn) {
	err := conn.Close()
	if err != nil {
		log.Errori(err)
	}
}

// TcpServer provides RECV-HANDLE-SENDBACK routine of TCP server module,
// Keepalive message must be handled carefully, because all communication
// data packets are walk on the single connection instance.
// TcpServer will be block running.
func TcpServer(host string, port int, proto TcpProtocol) {
	log.Warningfi("TCP server listen %s:%d", host, port)
	gtcp.NewServer(host+":"+strconv.Itoa(port), func(conn *gtcp.Conn) {
		log.Infoi("New connection ", conn.RemoteAddr().String())
		defer connectionClose(conn)
		p := proto.Instance()
		for {
			data, err := p.HandleMessage(conn)
			if err != nil {
				log.Errori(err)
				break
			}
			err = conn.Send(data)
			if err != nil {
				log.Errori(err)
				break
			}
		}
		log.Errori(conn.RemoteAddr().String(), " connection closed. ")
	}).Run()
}
