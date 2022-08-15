package db

import (
	"strings"
	"time"

	"github.com/lovelacelee/clsgo/pkg/log"
	"github.com/lovelacelee/clsgo/pkg/utils"
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

var dbsupported = []string{"MYSQL", "SQLITE", "SQLSERVER", "POSTGRES", "CLICKHOUSE"}

func ConnPoolSetting(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}

func (db *Db) open(dsn string) *gorm.DB {
	_db, err := gorm.Open(GormDriver[strings.ToUpper(db.Config.Type)](dsn), &gorm.Config{})
	if err != nil {
		log.Errori("failed to connect database")
	}
	ConnPoolSetting(_db)
	db.Orm = _db
	return _db
}

func (db *Db) Close() {
	sqlDB, err := db.Orm.DB()
	if err != nil {
		log.Errori(err)
	}
	if sqlDB != nil {
		sqlDB.Close()
	}
}

// New return the instance of open gorm instance,
// group was config name under database section.
// Returns nil, if group not found.
func New(group ...string) *Db {
	c := DbConfig{
		Type: "mysql",
	}
	g := "default"
	if !utils.IsEmpty(group) {
		g = group[0]
	}
	c = GetConfig(g)
	db := Db{Config: c}
	if !utils.CaseFoldIn(c.Type, dbsupported) {
		log.Errorfi("Unsupported database %s\n", c.Type)
		return nil
	}
	if utils.IsEmpty(c.Dsn) {
		log.Errorfi("Error database dns: %s\n", c.Dsn)
		return nil
	}
	log.Info(db.Config.Dsn)
	db.open(db.Config.Dsn)
	return &db
}
