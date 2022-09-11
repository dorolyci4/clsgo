package main

import (
	"context"
	"time"

	"github.com/lovelacelee/clsgo/pkg"
	"github.com/lovelacelee/clsgo/pkg/http"
	"github.com/lovelacelee/clsgo/pkg/log"
	"github.com/lovelacelee/clsgo/pkg/net"
	"github.com/lovelacelee/clsgo/pkg/redis"
	"github.com/lovelacelee/clsgo/pkg/version"
)

var Cfg = clsgo.Cfg

func init() {
}

type VersionResource struct {
}

type VersionResourceReq struct {
	http.Meta `path:"/version" method:"get" sm:"Get server version"`
}
type VersionResourceRes struct {
	Version string `dc:"Reply sever version"`
}

func (VersionResource) Version(context.Context, *VersionResourceReq) (res *VersionResourceRes, err error) {
	log.Debugfi(`Server version: %+v`, version.Version)
	res = &VersionResourceRes{
		Version: version.Version,
	}
	return
}

func simpleHTTPServer() {
	log.Infoi("ClsGO application ", version.Version)
	log.Infoi(Cfg.Get("database.default.0.link"))

	// HTTP simple web server
	apis := make(http.APIS)
	apis["/hello"] = func(r *http.Request) {
		r.Response.Write("Hello World!")
	}
	apig := make(http.APIG)
	resourceHandle := http.ResourceHandle{}
	resourceHandle.MiddlewareCallback = http.MiddlewareDefault
	resourceHandle.Res = VersionResource{}
	apig["/api/v1"] = resourceHandle
	http.App("0.0.0.0", 8080, "v1", &apis, &apig)

}

func tcpServer() {
	net.TcpServer("0.0.0.0", 8081, &HMProtocol{})
}
func forever() {
	redis := redis.New("default")
	defer redis.Close()
	tcpClient, err := net.TcpClient("0.0.0.0:8081", &HMProtocol{})
	if err != nil {
		defer tcpClient.Close()
	}
	var logon bool = false
	var e error
	for {
		time.Sleep(time.Second * 3)
		redis.Do("HSET", "hash", "init", "1")
		redis.Do("HSET", "hash", "key", "v")
		if !logon {
			e = tcpClient.Proto.Login(tcpClient.Conn)
			logon = (e == nil)
			tcpClient.ReconnectIfError(e)
		}
		e = tcpClient.Proto.KeepAlive(tcpClient.Conn)
		tcpClient.ReconnectIfError(e)
	}
}
func App() {
	go simpleHTTPServer()
	go tcpServer()
	go forever()
}
