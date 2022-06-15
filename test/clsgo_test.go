package clsgo

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/lovelacelee/clsgo/pkg"
)

func TestClsgo(t *testing.T) {
	v := clsgo.Version
	want := "v1.0.0"
	if reflect.TypeOf(v) != reflect.TypeOf(want) {
		t.Errorf("Not passed\n")
	} else {
		fmt.Printf("CLSGO: %s\n", v)
	}
}
