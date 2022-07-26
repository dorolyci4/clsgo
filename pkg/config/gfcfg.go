// Config implement use goframe, config automatic search supported,
// Only implement Get.
package config

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

// GCfgGet retrieves and returns value by specified `pattern`.
// It returns all values of current Json object if `pattern` is given empty or string ".".
// It returns nil if no value found by `pattern`.
//
// It returns a default value specified by `def` if value for `pattern` is not found.
func GCfgGet(pattern string, def ...interface{}) (x *gvar.Var) {
	var ctx = gctx.New()
	result, err := gcfg.Instance().Get(ctx, pattern)
	if err != nil {
		if len(def) > 0 {
			return gvar.New(def[0])
		} else {
			return gvar.New(nil)
		}
	} else {
		return result
	}
}
