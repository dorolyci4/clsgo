// reuse provides Listen and Dial functions that set socket
// options in order to be able to reuse ports. You should only use this
// package if you know what SO_REUSEADDR and SO_REUSEPORT are.
//
// For example:
//
//  // listen on the same port.
//  l1, _ := Listen("tcp", "127.0.0.1:1234")
//  l2, _ := Listen("tcp", "127.0.0.1:1234")
//
//  // dial from the same port.
//  l1, _ := Listen("tcp", "127.0.0.1:1234")
//  l2, _ := Listen("tcp", "127.0.0.1:1235")
//  c, _ := Dial("tcp", "127.0.0.1:1234", "127.0.0.1:1235")
//
// Note: can't dial self because tcp/ip stacks use 4-tuples to identify connections,
// and doing so would clash.
package tcp

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"time"
)

var (
	Enabled      = false
	listenConfig = net.ListenConfig{
		Control: Control,
	}
)

// Listen listens at the given network and address. see net.Listen
// Returns a net.Listener created from a file discriptor for a socket
// with SO_REUSEPORT and SO_REUSEADDR option set.
func Listen(network, address string) (net.Listener, error) {
	return listenConfig.Listen(context.Background(), network, address)
}

func TlsListen(network, laddr string, config *tls.Config) (net.Listener, error) {
	if config == nil || len(config.Certificates) == 0 &&
		config.GetCertificate == nil && config.GetConfigForClient == nil {
		return nil, errors.New("tls: neither Certificates, GetCertificate, nor GetConfigForClient set in Config")
	}
	l, err := Listen(network, laddr)
	if err != nil {
		return nil, err
	}
	return tls.NewListener(l, config), nil
}

// ListenPacket listens at the given network and address. see net.ListenPacket
// Returns a net.Listener created from a file discriptor for a socket
// with SO_REUSEPORT and SO_REUSEADDR option set.
func ListenPacket(network, address string) (net.PacketConn, error) {
	return listenConfig.ListenPacket(context.Background(), network, address)
}

// Dial dials the given network and address. see net.Dialer.Dial
// Returns a net.Conn created from a file discriptor for a socket
// with SO_REUSEPORT and SO_REUSEADDR option set.
func DialWithLocal(network, laddr, raddr string) (net.Conn, error) {
	nla, err := ResolveAddr(network, laddr)
	if err != nil {
		return nil, err
	}
	d := net.Dialer{
		Control:   Control,
		LocalAddr: nla,
	}
	return d.Dial(network, raddr)
}

func Dial(network, raddr string) (net.Conn, error) {
	d := net.Dialer{
		Control: Control,
	}
	return d.Dial(network, raddr)
}

// DialTimeout acts like Dial but takes a timeout.
//
// The timeout includes name resolution, if required.
// When using TCP, and the host in the address parameter resolves to
// multiple IP addresses, the timeout is spread over each consecutive
// dial, such that each is given an appropriate fraction of the time
// to connect.
//
// See func Dial for a description of the network and address
// parameters.
func DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	d := net.Dialer{
		Timeout: timeout,
		Control: Control,
	}
	return d.Dial(network, address)
}

// ResolveAddr parses given parameters to net.Addr.
func ResolveAddr(network, address string) (net.Addr, error) {
	switch network {
	case "ip", "ip4", "ip6":
		return net.ResolveIPAddr(network, address)
	case "tcp", "tcp4", "tcp6":
		return net.ResolveTCPAddr(network, address)
	case "udp", "udp4", "udp6":
		return net.ResolveUDPAddr(network, address)
	case "unix", "unixgram", "unixpacket":
		return net.ResolveUnixAddr(network, address)
	default:
		return nil, net.UnknownNetworkError(network)
	}
}
