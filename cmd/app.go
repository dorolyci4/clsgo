/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-30 17:13:53
 * @LastEditTime    : 2022-07-01 18:24:09
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /cmd/app.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package main

import (
	clsgo "github.com/lovelacelee/clsgo/pkg"
	"github.com/lovelacelee/clsgo/pkg/http"
	"github.com/lovelacelee/clsgo/pkg/log"
)

func App() {
	log.Info("ClsGO application %v", clsgo.Version)

	// HTTP simple web server
	apis := make(http.APIS)
	apis["/"] = func(r *http.Request) {
		r.Response.Write("Hello World!")
	}
	http.App("0.0.0.0", 8080, 8081, &apis)

}
