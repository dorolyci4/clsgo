/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-07-15 15:15:30
 * @LastEditTime    : 2022-07-15 15:42:18
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /pkg/crypto/md5.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

type MD5 struct {
	Data []byte
	hash string
}

func (ctx *MD5) Upper() string {
	ctx.hash = Md5(ctx.Data)
	return strings.ToUpper(ctx.hash)
}

func (ctx *MD5) Upper16() string {
	ctx.hash = Md5_16(ctx.Data)
	return strings.ToUpper(ctx.hash)
}

func Md5(in []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(in)
	cipher := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipher)
}

func Md5_16(in []byte) string {
	s := Md5(in)
	return s[8:24]
}
