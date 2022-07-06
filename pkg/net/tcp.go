/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-01-14 09:03:01
 * @LastEditTime    : 2022-07-06 18:37:41
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /pkg/net/tcp.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package net

import (
	"fmt"
	"strconv"

	"github.com/gogf/gf/v2/net/gtcp"
)

func TcpServer(host string, port int, protocol interface{}) {
	gtcp.NewServer(host+":"+strconv.Itoa(port), func(conn *gtcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(-1)
			if len(data) > 0 {
				if err := conn.Send(append([]byte("> "), data...)); err != nil {
					fmt.Println(err)
				}
			}
			if err != nil {
				break
			}
		}
	}).Run()
}
