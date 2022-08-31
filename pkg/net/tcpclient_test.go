package net_test

import (
	"github.com/lovelacelee/clsgo/pkg/log"
	"github.com/lovelacelee/clsgo/pkg/net"
	"github.com/lovelacelee/clsgo/pkg/utils"
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
	Example()
}
