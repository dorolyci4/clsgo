package net_test

import (
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/v1/log"
	"github.com/lovelacelee/clsgo/v1/net"
	"github.com/lovelacelee/clsgo/v1/utils"
	"testing"
)

// Example
func Example() {
	c1, err := net.DomainTcpClient("lovelacelee.com:8080")
	utils.IfError(err)
	log.Info(c1)
	_, err = net.DomainTcpClient("lovelacelee.com")
	utils.IfError(err)
	c3, err := net.TcpClient("192.168.137.100:22")
	utils.IfError(err)
	log.Info(c3)
}

func Test(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Run("DomainTCPClient", func(to *testing.T) {
			c, err := net.DomainTcpClient("lovelacelee.com:8080")
			t.AssertEQ(err, nil)
			t.AssertNE(c, nil)
		})
		t.Run("TCPClient", func(to *testing.T) {
			c, err := net.TcpClient("192.168.137.100:22")
			t.AssertNE(err, nil)
			t.Assert(c, nil)
		})
	})
}
