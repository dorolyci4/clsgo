/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-15 16:39:44
 * @LastEditTime    : 2022-07-05 19:52:48
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /pkg/http/server.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 * Simple http web server support apis and static files both.
 * GoFrame used.
 */
package http

import (
	"github.com/gogf/gf"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/lovelacelee/clsgo/pkg/log"
)

var ClsLog = log.ClsLog()

type Request = ghttp.Request

func init() {
	ClsLog.Info(gf.VERSION)
}

type APIS map[string]interface{}

func App(host string, portApi int, portWeb int, apis *APIS) {
	sApi := g.Server("API")
	for k, v := range *apis {
		sApi.BindHandler(k, v)
	}
	sApi.SetAddr(host)
	sApi.SetPort(portApi)
	sApi.Start()

	sStatic := g.Server("Static")
	sStatic.SetPort(portWeb)
	sStatic.SetIndexFolder(true)
	sStatic.SetServerRoot("public")
	sStatic.Start()

	g.Wait()
}
