package database

import (
	"fmt"
	"strings"
	"time"

	"github.com/lovelacelee/clsgo/v1/utils"
	"gorm.io/gorm"
)

type Db struct {
	Orm    *gorm.DB
	Config DbConfig
}

// See more https://gorm.io/zh_CN/docs/models.html
// https://gorm.io/zh_CN/docs/models.html#embedded_struct
// Field inluded:
//   ID        uint           `gorm:"primaryKey"`
//   CreatedAt time.Time
//   UpdatedAt time.Time
//   DeletedAt gorm.DeletedAt `gorm:"index"`
type Model struct {
	gorm.Model
}

func ConnPoolSetting(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if sqlDB != nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}
}

func (db *Db) open(dsn string) (*gorm.DB, error) {
	if utils.IsEmpty(GormDriver[strings.ToUpper(db.Config.Type)]) {
		return nil, fmt.Errorf("unsupported database type: %s", db.Config.Type)
	}
	_db, err := gorm.Open(GormDriver[strings.ToUpper(db.Config.Type)](dsn),
		&gorm.Config{
			// Creating and caching precompiled statements while executing
			// any SQL improves the speed of subsequent calls
			PrepareStmt: true,
		})
	ConnPoolSetting(_db)
	db.Orm = _db
	if db.Valid() {
		db.Orm.Use(CachePlugin)
	}
	return _db, err
}

func (db *Db) Valid() bool {
	return db.Orm != nil
}

func (db *Db) Close() {
	if db.Valid() {
		if sqlDB, _ := db.Orm.DB(); sqlDB != nil {
			sqlDB.Close()
		}
		db.Orm = nil
	}
}

// NewFromConfig return the instance of open gorm instance,
// group was config name under database section.
// Returns nil, if group not found.
func NewFromConfig(group ...string) *Db {
	c := GetConfig(group...)
	return New(c.Type, c.Dsn)
}

func New(dbtype string, dsn string) *Db {
	dbcfg := DbConfig{
		Dsn:  dsn,
		Type: dbtype,
	}
	db := Db{Config: dbcfg}
	db.open(dsn)
	return &db
}
