/*
 * @Author          : Lovelace
 * @Github          : https://github.com/lovelacelee
 * @Date            : 2022-06-13 15:40:43
 * @LastEditTime    : 2022-06-30 16:56:22
 * @LastEditors     : Lovelace
 * @Description     :
 * @FilePath        : /test/clsgo_test.go
 * Copyright 2022 Lovelace, All Rights Reserved.
 *
 *
 */
package clsgo

import (
	"reflect"
	"testing"

	"github.com/lovelacelee/clsgo/pkg"
	"github.com/lovelacelee/clsgo/pkg/log"
)

func TestClsgo(t *testing.T) {
	v := clsgo.Version
	want := "v1.0.0"
	if reflect.TypeOf(v) != reflect.TypeOf(want) {
		t.Errorf("Not passed\n")
	} else {
		log.Info("CLSGO: %s", v)
	}

}
