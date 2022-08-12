package db

import (
	"strings"
	"time"

	"github.com/lovelacelee/clsgo/pkg/log"
	"gorm.io/gorm"
)

type GormDb struct {
	Instance *gorm.DB
	DbType   string
}

// See more https://gorm.io/zh_CN/docs/models.html
//   ID        uint           `gorm:"primaryKey"`
//   CreatedAt time.Time
//   UpdatedAt time.Time
//   DeletedAt gorm.DeletedAt `gorm:"index"`
type Model struct {
	gorm.Model
}

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
func (gdb *GormDb) Open(dsn string) *gorm.DB {
	_db, err := gorm.Open(GormDriver[strings.ToUpper(gdb.DbType)](dsn), &gorm.Config{})
	if err != nil {
		log.Errori("failed to connect database")
	}
	ConnPoolSetting(_db)
	gdb.Instance = _db
	return _db
}

func NewSqlite(dsn string) *GormDb {
	db := GormDb{DbType: "SQLITE"}
	db.Open(dsn)
	return &db
}

func NewMysql(dsn string) *GormDb {
	db := GormDb{DbType: "MYSQL"}
	db.Open(dsn)
	return &db
}

func NewSqlServer(dsn string) *GormDb {
	db := GormDb{DbType: "SQLSERVER"}
	db.Open(dsn)
	return &db
}

func NewPostgres(dsn string) *GormDb {
	db := GormDb{DbType: "POSTGRES"}
	db.Open(dsn)
	return &db
}

func NewClickHouse(dsn string) *GormDb {
	db := GormDb{DbType: "CLICKHOUSE"}
	db.Open(dsn)
	return &db
}
