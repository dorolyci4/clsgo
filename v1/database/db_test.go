package database_test

import (
	"os"
	"testing"
	"time"

	"encoding/json"
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/lovelacelee/clsgo/v1/crypto"
	"github.com/lovelacelee/clsgo/v1/database"
	"github.com/lovelacelee/clsgo/v1/log"
	"github.com/lovelacelee/clsgo/v1/redis"
	"github.com/lovelacelee/clsgo/v1/utils"
)

const testCfg = `
database:
  default:
    dsn: "lee:lovelace@tcp(192.168.137.100:3306)/test?charset=utf8&parseTime=True&loc=Local"
    type: "mysql"
    redis: "cache"
  test:
    dsn: "lee:lovelace@tcp(192.168.137.100:3306)/test"
    type: "mysql"
  sqlite:
    dsn: "file::memory:?cache=shared"
    type: "sqlite"
redis.default:
  Address: 127.0.0.1:6379
  Db: 0
  pass: lovelacelee
  idleTimeout: 600
`

func TestMain(m *testing.M) {
	log.Green("")
	utils.DeleteFiles(".", "/*.yaml$")
	os.WriteFile("config.yaml", []byte(testCfg), 0755)
	time.Sleep(time.Microsecond * 30)
	m.Run()
}

// Data Model
type User struct {
	database.Model
	Name string `gorm:"name"`
}

// User compose cache key which must be unique in the lifetime of application
func CacheFind(db *database.Db, rdb *redis.Client, cacheKeyPrefix string, dest interface{}, conds ...interface{}) {
	if !db.Valid() {
		return
	}
	key := crypto.Md5Any(conds)
	cacheKeyPrefix += "_" + key
	if rdb != nil {
		r, err := rdb.Do("EXISTS", cacheKeyPrefix)
		if err != nil || !r.Bool() {
			goto DBFIND
		} else {
			r, err := rdb.Do("GET", cacheKeyPrefix)
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
	if rdb != nil {
		// expired 2 seconds
		rdb.Do("SETEX", cacheKeyPrefix, 10, dest)
	}
}

func Test(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Run("Config", func(_ *testing.T) {
			cfg := database.GetConfig()
			t.Assert(cfg.Type, "mysql")
			utils.DeleteFiles(".", "/*.yaml$")
			cfg = database.GetConfig()
			t.Assert(cfg.Type, "sqlite")
		})
		t.Run("AutoMigrate", func(_ *testing.T) {
			utils.DeleteFiles(".", "/*.db$")
			db := database.NewFromConfig()
			if db != nil {
				defer db.Close()
				if db.Valid() {
					db.Orm.Use(database.TracePlugin)
					db.Orm.AutoMigrate(&User{})
					user := User{
						Name: "Lee",
					}
					db.Orm.Create(&user)
					quser := User{}
					var count int64 = 0
					db.Orm.Find(&quser, "id = ?", 2)
					db.Orm.Model(&User{}).Count(&count)
					t.Assert(quser.ID, 0)
					t.Assert(count, 1)
				}
			}
			utils.DeleteFiles(".", "/*.db$")
			database.New("somedb", "lee:lovelace@tcp(192.168.137.100:3306)/test")
			database.New("mysql", "lee:lovelace@tcp(192.168.137.100:3306)/test")
			database.New("sqlserver", "sqlserver://username:password@localhost:9930?database=gorm")
			database.New("postgresql", "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai")
			database.New("clickhouse", "tcp://localhost:9000?database=gorm&username=gorm&password=gorm&read_timeout=10&write_timeout=20")
		})
	})
}
