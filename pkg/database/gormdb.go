package database

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/lovelacelee/clsgo/pkg/log"
	"github.com/lovelacelee/clsgo/pkg/redis"
	"github.com/lovelacelee/clsgo/pkg/utils"

	"github.com/lovelacelee/clsgo/pkg/crypto"
	"gorm.io/gorm"
)

type Db struct {
	Orm    *gorm.DB
	Config DbConfig
	RDB    *redis.Client
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
	_db, err := gorm.Open(GormDriver[strings.ToUpper(db.Config.Type)](dsn), &gorm.Config{
		// Creating and caching precompiled statements while executing any SQL improves the speed of subsequent calls
		PrepareStmt: true,
	})
	if err != nil {
		log.Errorfi("failed to connect database: %v", err)
		return nil
	}
	ConnPoolSetting(_db)
	db.Orm = _db
	return _db
}

func (db *Db) Valid() bool {
	return db.Orm != nil
}

func (db *Db) Close() {
	if db.Valid() {
		sqlDB, _ := db.Orm.DB()
		if sqlDB != nil {
			sqlDB.Close()
		}
	}
	if db.RDB != nil {
		db.RDB.Close()
	}
}

// User compose cache key which must be unique in the lifetime of application
func (db *Db) CacheFind(cacheKeyPrefix string, dest interface{}, conds ...interface{}) {
	if !db.Valid() {
		return
	}
	key := crypto.Md5Any(conds)
	cacheKeyPrefix += "_" + key
	if db.RDB != nil {
		r, err := db.RDB.Do("EXISTS", cacheKeyPrefix)
		if err != nil || !r.Bool() {
			goto DBFIND
		} else {
			r, err := db.RDB.Do("GET", cacheKeyPrefix)
			if err != nil {
				goto DBFIND
			}
			json.Unmarshal(r.Bytes(), dest)
			return
		}
	}
DBFIND:
	db.Orm.Find(dest, conds...)
	// Write cache
	if db.RDB != nil {
		// expired 2 seconds
		db.RDB.Do("SETEX", cacheKeyPrefix, 10, dest)
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
	// log.Infoi(c)
	db := Db{Config: c}
	if !utils.CaseFoldIn(c.Type, dbsupported) {
		log.Errorfi("Unsupported database %s\n", c.Type)
		return nil
	}
	if utils.IsEmpty(c.Dsn) {
		log.Errorfi("Error database dns: %s\n", c.Dsn)
		return nil
	}
	if !utils.IsEmpty(db.Config.Redis) {
		db.RDB = redis.New("default")
	}
	db.open(db.Config.Dsn)
	return &db
}
