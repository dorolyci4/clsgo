/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-30 17:13:53
 * @LastEditTime    : 2022-07-07 19:08:57
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /cmd/app.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"time"

	"github.com/gogf/gf/v2/net/gtcp"
	clsgo "github.com/lovelacelee/clsgo/pkg"
	"github.com/lovelacelee/clsgo/pkg/config"
	"github.com/lovelacelee/clsgo/pkg/http"
	"github.com/lovelacelee/clsgo/pkg/log"
	"github.com/lovelacelee/clsgo/pkg/net"
)

var l = log.ClsLog()

func App() {
	l.Info("ClsGO application ", clsgo.Version)
	l.Info(config.Get("database.default.0.link"))

	// HTTP simple web server
	apis := make(http.APIS)
	apis["/"] = func(r *http.Request) {
		r.Response.Write("Home")
	}
	apis["/hello"] = func(r *http.Request) {
		r.Response.Write("Hello World!")
	}
	http.App("0.0.0.0", 8080, "v1", &apis)

}

var (
	ErrOk            = errors.New("ok")
	ErrBadCredential = errors.New("bad credentials, username or password invalid")
)

type HMHeadV1 struct {
	ID  uint32
	Len uint32
	Err uint32
}

type HMHead struct {
	V1  HMHeadV1
	Sid uint32
}

type HMClient struct {
	SN            string
	IP            string
	Port          uint16
	LastKeepAlive time.Time
	Uptime        time.Duration
	BootTime      time.Time
	Authorized    bool
}

type HMProtocol struct {
	Head HMHead
	Body []byte
	Cli  HMClient
}

func (p *HMProtocol) onLogin(conn *net.Conn) error {
	data, err := conn.Recv(int(p.Head.V1.Len), gtcp.Retry{Count: 3, Interval: 10})
	if err != nil {
		return err
	}
	log.Info(string(data))
	p.Cli.Authorized = true
	return nil
}

func (p *HMProtocol) OnHead(conn *net.Conn) error {
	var err error = nil
	l.Infof("0x%08X", p.Head.V1.ID)
	switch p.Head.V1.ID {
	case 0x0000060D:
		err = p.onLogin(conn)
		break
	default:
		break
	}
	return err
}

func (p *HMProtocol) OnBody(conn *net.Conn) error {
	return nil

}

func (p *HMProtocol) Send(conn *net.Conn) error {
	return nil

}

func (p *HMProtocol) Recv(conn *net.Conn) error {
	// First connected, authorization required
	if !p.Cli.Authorized {
		data, err := conn.Recv(12, gtcp.Retry{Count: 3, Interval: 10})
		if err != nil {
			return err
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.BigEndian, &p.Head.V1); err != nil {
			return err
		}
		return p.OnHead(conn)
	} else {
		data, err := conn.Recv(16, gtcp.Retry{Count: 3, Interval: 10})
		if err != nil {
			return err
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.BigEndian, &p.Head); err != nil {
			return err
		}
		return p.OnHead(conn)
	}
}

func TCPApp() {
	var proto net.TcpProtocol = &HMProtocol{}

	net.TcpServer("0.0.0.0", 8081, &proto)
}
