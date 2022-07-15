/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-07-15 15:07:16
 * @LastEditTime    : 2022-07-15 15:09:01
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /pkg/utils/id.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func SessionId() string {
	buf := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(buf)
}
