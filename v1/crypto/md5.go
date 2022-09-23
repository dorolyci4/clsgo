package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
	"strings"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/lovelacelee/clsgo/v1/utils"
)

type MD5 struct {
	hash hash.Hash //md5
}

const MD5_16 = true

func (ctx *MD5) Continue(s string) *MD5 {
	ctx.hash.Write([]byte(s))
	return ctx
}

func (ctx *MD5) Sum(b16ornot ...bool) string {
	b16 := utils.Param(b16ornot, false)
	cipher := ctx.hash.Sum(nil)
	sumstring := hex.EncodeToString(cipher)
	if b16 {
		return sumstring[8:24]
	}
	return sumstring
}

func (ctx *MD5) SumUpper(b16ornot ...bool) string {
	return strings.ToUpper(ctx.Sum(b16ornot...))
}

func NewMD5(s ...string) *MD5 {
	ctx := &MD5{
		hash: md5.New(),
	}
	for _, v := range s {
		ctx.Continue(v)
	}
	return ctx
}

// Return lower md5sum string
func MD5Sum(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[0:])
}

func Md5Any(in ...any) string {
	md5Ctx := md5.New()
	for _, v := range in {
		md5Ctx.Write(gconv.Bytes(v))
	}
	cipher := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipher)
}
