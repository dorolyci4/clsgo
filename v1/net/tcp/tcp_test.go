// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package tcp_test

import (
	"fmt"
	"github.com/lovelacelee/clsgo/v1/net/tcp"
)

func ExampleGetFreePort() {
	fmt.Println(tcp.GetFreePort())

	// May Output:
	// 57429 <nil>
}

func ExampleGetFreePorts() {
	fmt.Println(tcp.GetFreePorts(2))

	// May Output:
	// [57743 57744] <nil>
}
