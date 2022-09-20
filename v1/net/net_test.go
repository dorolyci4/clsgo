package net_test

import (
	"sync"
	"testing"
	"time"

	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/v1/log"
	"github.com/lovelacelee/clsgo/v1/net"
	"github.com/lovelacelee/clsgo/v1/utils"
)

const testServerName = "test-tcp-server"
const testMessage = "test\n"

type TestProtocol struct {
}

func (p *TestProtocol) ServerName() string {
	return testServerName
}

// Method of handling message
func (p *TestProtocol) HandleMessage(conn *net.Conn) ([]byte, error) {
	data, err := conn.RecvLine()
	if !utils.IsEmpty(data) {
		log.Info(data)
	}
	time.Sleep(time.Second * 3)
	return []byte("response message"), err
}

func (p *TestProtocol) Instance() net.TcpProtocol {
	return p
}

// Example
func Example() {
	// c1, err := net.DomainTcpClient("lovelacelee.com:8080")
	// utils.IfError(err)
	// log.Info(c1)
	// _, err = net.DomainTcpClient("lovelacelee.com")
	// utils.IfError(err)
	// c3, err := net.TcpClient("192.168.137.100:22")
	// utils.IfError(err)
	// log.Info(c3)
}

type ServerParam struct {
	Domain   string
	CliIsNil bool
	ErrorNil bool
}

var ServerParamCases = []ServerParam{
	{"localhost", true, false},
	{"localhost:19090", false, true},
	{"localhos:19090", true, false},
}

func TestNetPackage(t *testing.T) {
	s := net.NewTcpServer("0.0.0.0", 19090, &TestProtocol{})
	defer s.Close()

	var wg sync.WaitGroup
	wg.Add(4)
	gtest.C(t, func(t *gtest.T) {
		t.Run("TcpServer", func(_ *testing.T) {
			t.AssertNE(s, nil)
			go s.Run()
			wg.Done()
		})
		t.Run("DomainTCPClient", func(_ *testing.T) {
			for _, casei := range ServerParamCases {
				c, err := net.DomainTcpClient(casei.Domain)
				t.Assert(err == nil, casei.ErrorNil)
				t.Assert(c == nil, casei.CliIsNil)
			}
			c, err := net.DomainTcpClient("localhost:19090")
			t.AssertEQ(err, nil)
			t.AssertNE(c, nil)
			c.Close()
			c.Reconnect()
			c.Close()
			c.ReconnectIfError(nil)
			c.ReconnectIfError(utils.ErrTcpDomain)
			defer c.Close()

			wg.Done()
		})
		t.Run("ClientClose", func(_ *testing.T) {
			c, err := net.DomainTcpClient("localhost:19090", &TestProtocol{})
			t.AssertEQ(err, nil)
			t.AssertNE(c, nil)

			t.Assert(c.Conn.Send([]byte(testMessage)), nil)
			c.Close()
			wg.Done()
		})
		t.Run("TCPClient", func(_ *testing.T) {
			c, err := net.TcpClient("192.168.88.100:22") //Assume this ip:port cannot be connected
			t.AssertNE(err, nil)
			t.Assert(c, nil)

			wg.Done()
		})
	})
	wg.Wait()
}
