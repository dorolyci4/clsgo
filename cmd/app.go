package main

import (
	"github.com/lovelacelee/clsgo/pkg"
	"github.com/lovelacelee/clsgo/pkg/config"
	"github.com/lovelacelee/clsgo/pkg/http"
	"github.com/lovelacelee/clsgo/pkg/log"
	"github.com/lovelacelee/clsgo/pkg/net"
)

var Cfg = clsgo.Cfg

func init() {
	// log.LogInfo("clsgo application")
}

func simpleHTTPServer() {
	log.Info("ClsGO application ", clsgo.Version)
	log.Info(config.Get("database.default.0.link"))

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

func tcpServer() {
	net.TcpServer("0.0.0.0", 8081, &HMProtocol{})
}

func App() {
	go simpleHTTPServer()
	go tcpServer()
}
