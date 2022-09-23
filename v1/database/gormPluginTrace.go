package database

import (
	"time"

	"github.com/lovelacelee/clsgo/v1/log"
	"gorm.io/gorm"
)

// Plugin trace sql and time of it excuted
type PluginTrace struct{}

func (op *PluginTrace) Name() string {
	return "PluginTrace"
}

func (op *PluginTrace) Initialize(db *gorm.DB) (err error) {
	// Before
	_ = db.Callback().Create().Before("gorm:create").Register("trace-before-create", tracePluginBefore)
	_ = db.Callback().Query().Before("gorm:query").Register("trace-before-query", tracePluginBefore)
	_ = db.Callback().Delete().Before("gorm:delete").Register("trace-before-delete", tracePluginBefore)
	_ = db.Callback().Update().Before("gorm:update").Register("trace-before-update", tracePluginBefore)
	_ = db.Callback().Row().Before("gorm:row").Register("trace-before-row", tracePluginBefore)
	_ = db.Callback().Raw().Before("gorm:raw").Register("trace-before-raw", tracePluginBefore)

	// After
	_ = db.Callback().Create().After("gorm:create").Register("trace-after-create", tracePluginAfter)
	_ = db.Callback().Query().After("gorm:query").Register("trace-after-query", tracePluginAfter)
	_ = db.Callback().Delete().After("gorm:delete").Register("trace-after-delete", tracePluginAfter)
	_ = db.Callback().Update().After("gorm:update").Register("trace-after-update", tracePluginAfter)
	_ = db.Callback().Row().After("gorm:row").Register("trace-after-row", tracePluginAfter)
	_ = db.Callback().Raw().After("gorm:raw").Register("trace-after-raw", tracePluginAfter)
	return
}

var TracePlugin gorm.Plugin = &PluginTrace{}

func tracePluginBefore(db *gorm.DB) {
	db.InstanceSet("startTime", time.Now())
}

func tracePluginAfter(db *gorm.DB) {
	_ts, isExist := db.InstanceGet("startTime")
	if isExist {
		ts, ok := _ts.(time.Time)
		if ok {
			sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
			log.Debugfi("%s affected %d rows cost %f seconds", sql, db.Statement.RowsAffected, time.Since(ts).Seconds())
		}
	}
}
