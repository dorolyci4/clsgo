/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-30 17:13:53
 * @LastEditTime    : 2022-06-30 17:27:15
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /cmd/app.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package main

import (
	"fmt"

	clsgo "github.com/lovelacelee/clsgo/pkg"
	// "github.com/lovelacelee/clsgo/pkg/utils"
)

func App() {
	fmt.Printf("ClsGO application %v!\n", clsgo.Version)
}
