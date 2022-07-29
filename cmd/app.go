package main

import (
	"context"

	"github.com/lovelacelee/clsgo/pkg"
	"github.com/lovelacelee/clsgo/pkg/http"
	"github.com/lovelacelee/clsgo/pkg/log"
	"github.com/lovelacelee/clsgo/pkg/net"
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
	log.Debugf(`Server version: %+v`, clsgo.Version)
	res = &VersionResourceRes{
		Version: clsgo.Version,
	}
	return
}

func simpleHTTPServer() {
	log.Info("ClsGO application ", clsgo.Version)
	log.Info(Cfg.Get("database.default.0.link"))

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

func App() {
	go simpleHTTPServer()
	go tcpServer()
}
