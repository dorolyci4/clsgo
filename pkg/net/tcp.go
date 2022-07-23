package net

import (
	"strconv"

	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/lovelacelee/clsgo/pkg/log"
)

type Conn = gtcp.Conn
type Retry = gtcp.Retry

var l = log.ClsLog(&log.ClsFormatter{
	Prefix:      true,
	ForceColors: false,
})

type TcpProtocol interface {
	HandleMessage(conn *Conn) ([]byte, error)
	Instance() TcpProtocol
}

func connectionClose(conn *gtcp.Conn) {
	err := conn.Close()
	if err != nil {
		l.Error(err)
	}
}

func TcpServer(host string, port int, proto TcpProtocol) {
	l.Warningf("TCP server listen %s:%d", host, port)
	gtcp.NewServer(host+":"+strconv.Itoa(port), func(conn *gtcp.Conn) {
		l.Info("New connection ", conn.RemoteAddr().String())
		defer connectionClose(conn)
		p := proto.Instance()
		for {
			data, err := p.HandleMessage(conn)
			if err != nil {
				l.Error(err)
				break
			}
			err = conn.Send(data)
			if err != nil {
				l.Error(err)
				break
			}
		}
		l.Error(conn.RemoteAddr().String(), " connection closed. ")
	}).Run()
}
