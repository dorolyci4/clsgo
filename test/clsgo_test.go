// Test package for clsgo
package clsgo_test

import (
	"reflect"
	"testing"

	"github.com/lovelacelee/clsgo/pkg"
	"github.com/lovelacelee/clsgo/pkg/crypto"
	"github.com/lovelacelee/clsgo/pkg/log"
	"github.com/lovelacelee/clsgo/pkg/utils"
)

func TestClsgo(t *testing.T) {
	v := clsgo.Version
	want := "v1.0.0"
	if reflect.TypeOf(v) != reflect.TypeOf(want) {
		t.Errorf("Not passed\n")
	} else {
		log.Infof("CLSGO: %s", v)
	}
	log.Info(utils.SessionId())
	log.Info(crypto.Md5([]byte(want)))
	log.Info(crypto.Md5_16([]byte(want)))
	md5 := crypto.MD5{Data: []byte(want)}
	log.Info(md5.Upper())
	log.Info(md5.Upper16())
}
