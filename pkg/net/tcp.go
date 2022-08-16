// Package net provides basic network functions,
// TCP/UDP/WebSocket quick implements.
package net

import (
	"strconv"

	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/lovelacelee/clsgo/pkg/log"
)

type Conn = gtcp.Conn
type Retry = gtcp.Retry

type TcpProtocol interface {
	HandleMessage(conn *Conn) ([]byte, error)
	Instance() TcpProtocol
}

func connectionClose(conn *gtcp.Conn) {
	err := conn.Close()
	if err != nil {
		log.Errori(err)
	}
}

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
