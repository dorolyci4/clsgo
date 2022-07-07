/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-07-07 18:25:03
 * @LastEditTime    : 2022-07-07 18:28:35
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /pkg/utils/error.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */

package utils

import (
	"errors"
)

var (
	ErrEOF           = errors.New("EOF")
	ErrUnexpectedEOF = errors.New("unexpected EOF")
	ErrNoProgress    = errors.New("multiple Read calls return no data or error")
)
