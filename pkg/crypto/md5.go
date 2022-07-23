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
