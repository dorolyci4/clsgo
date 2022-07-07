/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-01-14 09:03:01
 * @LastEditTime    : 2022-07-07 19:03:38
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /pkg/net/tcp.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package net

import (
	"strconv"

	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/lovelacelee/clsgo/pkg/log"
)

type Conn = gtcp.Conn

var l = log.ClsLog()

type TcpProtocol interface {
	OnHead(conn *Conn) error
	OnBody(conn *Conn) error
	Recv(conn *Conn) error
	Send(conn *Conn) error
}

func TcpServer(host string, port int, proto *TcpProtocol) {
	l.Warningf("TCP server listen %s:%d", host, port)
	gtcp.NewServer(host+":"+strconv.Itoa(port), func(conn *gtcp.Conn) {
		defer conn.Close()
		for {
			err := (*proto).Recv(conn)
			if err != nil {
				l.Error(err)
				break
			}
			err = (*proto).Send(conn)
			if err != nil {
				l.Error(err)
				break
			}
		}
	}).Run()
}
