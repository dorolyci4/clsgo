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
	_ = db.Callback().Create().Before("gorm:before_create").Register("callbackBefore", before)
	_ = db.Callback().Query().Before("gorm:query").Register("callbackBefore", before)
	_ = db.Callback().Delete().Before("gorm:before_delete").Register("callbackBefore", before)
	_ = db.Callback().Update().Before("gorm:setup_reflect_value").Register("callbackBefore", before)
	_ = db.Callback().Row().Before("gorm:row").Register("callbackBefore", before)
	_ = db.Callback().Raw().Before("gorm:raw").Register("callbackBefore", before)

	// After
	_ = db.Callback().Create().After("gorm:after_create").Register("callbackAfter", after)
	_ = db.Callback().Query().After("gorm:after_query").Register("callbackAfter", after)
	_ = db.Callback().Delete().After("gorm:after_delete").Register("callbackAfter", after)
	_ = db.Callback().Update().After("gorm:after_update").Register("callbackAfter", after)
	_ = db.Callback().Row().After("gorm:row").Register("callbackAfter", after)
	_ = db.Callback().Raw().After("gorm:raw").Register("callbackAfter", after)
	return
}

var TracePlugin gorm.Plugin = &PluginTrace{}

func before(db *gorm.DB) {
	db.InstanceSet("startTime", time.Now())
}

func after(db *gorm.DB) {
	_ts, isExist := db.InstanceGet("startTime")
	if !isExist {
		return
	}
	ts, ok := _ts.(time.Time)
	if !ok {
		return
	}
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	log.Infofi("%s affected %d rows cost %f seconds", sql, db.Statement.RowsAffected, time.Since(ts).Seconds())
}
