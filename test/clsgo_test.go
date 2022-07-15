/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-13 15:40:43
 * @LastEditTime    : 2022-07-15 15:42:34
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /test/clsgo_test.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package clsgo_test

import (
	"reflect"
	"testing"

	"github.com/lovelacelee/clsgo/pkg"
	"github.com/lovelacelee/clsgo/pkg/crypto"
	"github.com/lovelacelee/clsgo/pkg/log"
	"github.com/lovelacelee/clsgo/pkg/utils"
)

var l = log.ClsLog()

func TestClsgo(t *testing.T) {
	v := clsgo.Version
	want := "v1.0.0"
	if reflect.TypeOf(v) != reflect.TypeOf(want) {
		t.Errorf("Not passed\n")
	} else {
		l.Infof("CLSGO: %s", v)
	}
	l.Info(utils.SessionId())
	l.Info(crypto.Md5([]byte(want)))
	l.Info(crypto.Md5_16([]byte(want)))
	md5 := crypto.MD5{Data: []byte(want)}
	l.Info(md5.Upper())
	l.Info(md5.Upper16())
}
