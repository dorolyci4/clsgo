package net

import (
	"strings"
	"time"

	"github.com/gogf/gf/v2/net/gipv4"
	"github.com/gogf/gf/v2/net/gtcp"
	"github.com/lovelacelee/clsgo/v1/config"
	"github.com/lovelacelee/clsgo/v1/utils"
)

type Client struct {
	timeout  time.Duration
	ipServer string
	Conn     *gtcp.Conn
	Proto    TcpProtocol
}

// Establish a TCP connection to "server:port" with specified protocol(proto),
// Config: "server.tcpTimeout" unit:seconds
func TcpClient(ipserver string, proto ...TcpProtocol) (*Client, error) {
	timeout := config.GetDurationWithDefault("server.tcpTimeout", 5)
	c, err := gtcp.NewConn(ipserver, timeout*time.Second)
	if err != nil || c == nil {
		return nil, err
	}
	client := Client{Conn: c, ipServer: ipserver, timeout: timeout}
	if !utils.IsEmpty(proto) {
		client.Proto = proto[0]
	}
	return &client, nil
}

// Establish a TCP connection to "domain:port" with specified protocol(proto),
// Config: "server.tcpTimeout" unit:seconds
func DomainTcpClient(domainServer string, proto ...TcpProtocol) (*Client, error) {
	host := strings.Split(domainServer, ":")
	if len(host) < 2 {
		return nil, utils.ErrTcpDomain
	}
	s, e := gipv4.GetHostByName(host[0])
	if e != nil {
		return nil, e
	}
	return TcpClient(s+":"+host[1], proto...)
}

func (client *Client) Close() {
	if client.Conn != nil {
		client.Conn.Close()
	}
	client.Conn = nil
}

func (client *Client) Reconnect() error {
	client.Close()
	c, err := gtcp.NewConn(client.ipServer, client.timeout*time.Second)
	client.Conn = c
	return err
}

func (client *Client) ReconnectIfError(e error) {
	if e != nil {
		client.Reconnect()
	}
}
