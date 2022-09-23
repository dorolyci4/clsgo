// Default internal cache plugin for gorm
// See more on https://gorm.io/zh_CN/docs/write_plugins.html

package database

import (
	// "github.com/gogf/gf/v2/os/gcache"
	// "github.com/gogf/gf/v2/os/gctx"
	// "github.com/lovelacelee/clsgo/v1/log"
	"gorm.io/gorm"
)

// Plugin trace sql and time of it excuted
type PluginCache struct{}

var (
// cachectx = gctx.New()
// cache    = gcache.New()
)

func (op *PluginCache) Name() string {
	return "PluginCache"
}

func (op *PluginCache) Initialize(db *gorm.DB) (err error) {
	// Before
	_ = db.Callback().Create().Before("gorm:create").Register("cache-before-create", cachePluginBefore)
	_ = db.Callback().Query().Before("gorm:query").Register("cache-before-query", cachePluginBefore)
	_ = db.Callback().Delete().Before("gorm:delete").Register("cache-before-delete", cachePluginBefore)
	_ = db.Callback().Update().Before("gorm:update").Register("cache-before-update", cachePluginBefore)
	_ = db.Callback().Row().Before("gorm:row").Register("cache-before-row", cachePluginBefore)
	_ = db.Callback().Raw().Before("gorm:raw").Register("cache-before-raw", cachePluginBefore)

	// After
	_ = db.Callback().Create().After("gorm:create").Register("cache-after-create", cachePluginAfter)
	_ = db.Callback().Query().After("gorm:query").Register("cache-after-query", cachePluginAfter)
	_ = db.Callback().Delete().After("gorm:delete").Register("cache-after-delete", cachePluginAfter)
	_ = db.Callback().Update().After("gorm:udpate").Register("cache-after-update", cachePluginAfter)
	_ = db.Callback().Row().After("gorm:row").Register("cache-after-row", cachePluginAfter)
	_ = db.Callback().Raw().After("gorm:raw").Register("cache-after-raw", cachePluginAfter)
	return
}

var CachePlugin gorm.Plugin = &PluginCache{}

func cachePluginBefore(db *gorm.DB) {
}

func cachePluginAfter(db *gorm.DB) {
	// sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
}
